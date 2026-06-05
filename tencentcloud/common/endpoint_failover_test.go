package common

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
)

// =====================================================================
// Pure host classification
// =====================================================================

func Test_tldMatchOf_plainCom(t *testing.T) {
	m := tldMatchOf("cvm.tencentcloudapi.com")
	if m == nil {
		t.Fatal("expected non-nil match")
	}
	if m.familyIdx != 0 || m.tldIdx != 0 || m.hasRegion || m.servicePrefix != "cvm" {
		t.Fatalf("got %+v, want familyIdx=0 tldIdx=0 !hasRegion service=cvm", m)
	}
}

func Test_tldMatchOf_plainCn(t *testing.T) {
	m := tldMatchOf("cvm.tencentcloudapi.cn")
	if m == nil || m.tldIdx != 1 {
		t.Fatalf("got %+v, want tldIdx=1", m)
	}
}

func Test_tldMatchOf_plainComCn(t *testing.T) {
	m := tldMatchOf("cvm.tencentcloudapi.com.cn")
	if m == nil || m.tldIdx != 2 {
		t.Fatalf("got %+v, want tldIdx=2", m)
	}
}

func Test_tldMatchOf_aiFamily(t *testing.T) {
	m := tldMatchOf("hunyuan.ai.tencentcloudapi.com")
	if m == nil || m.familyIdx != 1 || m.tldIdx != 0 || m.servicePrefix != "hunyuan" {
		t.Fatalf("got %+v, want familyIdx=1 tldIdx=0 service=hunyuan", m)
	}
}

func Test_tldMatchOf_internalFamily(t *testing.T) {
	m := tldMatchOf("cvm.internal.tencentcloudapi.com")
	if m == nil || m.familyIdx != 2 || m.servicePrefix != "cvm" {
		t.Fatalf("got %+v, want familyIdx=2 service=cvm", m)
	}
}

func Test_tldMatchOf_regionPinned(t *testing.T) {
	m := tldMatchOf("cvm.ap-guangzhou.tencentcloudapi.com")
	if m == nil || !m.hasRegion {
		t.Fatalf("got %+v, want hasRegion=true", m)
	}
}

func Test_tldMatchOf_unknownHost(t *testing.T) {
	if m := tldMatchOf("example.com"); m != nil {
		t.Fatalf("expected nil for example.com, got %+v", m)
	}
}

func Test_tldMatchOf_emptyHost(t *testing.T) {
	if m := tldMatchOf(""); m != nil {
		t.Fatalf("expected nil for empty host, got %+v", m)
	}
}

func Test_tldMatchOf_woaSuffixIsUnknown(t *testing.T) {
	if m := tldMatchOf("cvm.tencentcloudapi.woa.com"); m != nil {
		t.Fatalf("expected nil for non-Tencent suffix, got %+v", m)
	}
}

// =====================================================================
// Plan generation
// =====================================================================

func newFailoverEnabledClient() *Client {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	return NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
}

func Test_candidatesFor_backupEndpointTwoCandidates(t *testing.T) {
	c := newFailoverEnabledClient()
	c.profile.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"

	// region-pinned host should still get candidates when backupEndpoint is set
	cs := candidatesFor(c, "cvm.ap-shanghai.tencentcloudapi.com")
	if cs == nil {
		t.Fatal("expected candidates for region-pinned host with backupEndpoint")
	}
	if len(cs) != 2 {
		t.Fatalf("got %d candidates, want 2", len(cs))
	}
	if cs[0].host != "cvm.ap-shanghai.tencentcloudapi.com" {
		t.Fatalf("candidate[0] = %s, want origin", cs[0].host)
	}
	if cs[1].host != "cvm.ap-guangzhou.tencentcloudapi.com" {
		t.Fatalf("candidate[1] = %s, want backup-prefixed host", cs[1].host)
	}
}

func Test_candidatesFor_plainHostThreeCandidates(t *testing.T) {
	c := newFailoverEnabledClient()
	cs := candidatesFor(c, "cvm.tencentcloudapi.com")
	if len(cs) != 3 {
		t.Fatalf("got %d candidates, want 3", len(cs))
	}
	got := []string{cs[0].host, cs[1].host, cs[2].host}
	want := []string{
		"cvm.tencentcloudapi.com",
		"cvm.tencentcloudapi.cn",
		"cvm.tencentcloudapi.com.cn",
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("candidate[%d] = %s, want %s", i, got[i], want[i])
		}
	}
}

func Test_candidatesFor_aiFamilyStaysInFamily(t *testing.T) {
	c := newFailoverEnabledClient()
	cs := candidatesFor(c, "hunyuan.ai.tencentcloudapi.com")
	if len(cs) != 3 {
		t.Fatalf("got %d candidates", len(cs))
	}
	for _, cand := range cs {
		if !strings.Contains(cand.host, ".ai.") {
			t.Fatalf("candidate %s leaked outside ai. family", cand.host)
		}
	}
}

func Test_candidatesFor_internalFamilyStaysInFamily(t *testing.T) {
	c := newFailoverEnabledClient()
	cs := candidatesFor(c, "cvm.internal.tencentcloudapi.com")
	if len(cs) != 3 {
		t.Fatalf("got %d candidates", len(cs))
	}
	for _, cand := range cs {
		if !strings.Contains(cand.host, ".internal.") {
			t.Fatalf("candidate %s leaked outside internal. family", cand.host)
		}
	}
}

func Test_candidatesFor_regionPinnedNoBackup_isNil(t *testing.T) {
	c := newFailoverEnabledClient()
	if cs := candidatesFor(c, "cvm.ap-guangzhou.tencentcloudapi.com"); cs != nil {
		t.Fatalf("region-pinned host without backup should be pass-through, got %+v", cs)
	}
}

func Test_candidatesFor_unknownHost_isNil(t *testing.T) {
	c := newFailoverEnabledClient()
	if cs := candidatesFor(c, "example.com"); cs != nil {
		t.Fatalf("unknown host should be pass-through, got %+v", cs)
	}
}

// =====================================================================
// End-to-end with mock RoundTripper
// =====================================================================

const failoverSuccessResp = `{"Response": {"RequestId": "req-ok"}}`

// scriptedRT is a programmable RoundTripper. Each call pops the next
// outcome from the queue and either returns a programmed response or fails
// with a programmed error. Records every request's URL host and the
// Authorization header for inspection.
type scriptedRT struct {
	mu        sync.Mutex
	outcomes  []rtOutcome
	hostsSeen []string
	authsSeen []string
}

type rtOutcome struct {
	err  error
	code int
	body string
}

func (r *scriptedRT) programOk() {
	r.outcomes = append(r.outcomes, rtOutcome{code: 200, body: failoverSuccessResp})
}

func (r *scriptedRT) programResponse(code int, body string) {
	r.outcomes = append(r.outcomes, rtOutcome{code: code, body: body})
}

func (r *scriptedRT) programErr(err error) {
	r.outcomes = append(r.outcomes, rtOutcome{err: err})
}

func (r *scriptedRT) RoundTrip(req *http.Request) (*http.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	host := req.Host
	if host == "" {
		host = req.URL.Host
	}
	r.hostsSeen = append(r.hostsSeen, host)
	r.authsSeen = append(r.authsSeen, req.Header.Get("Authorization"))

	if len(r.outcomes) == 0 {
		return nil, errors.New("scriptedRT exhausted: no programmed outcome left")
	}
	o := r.outcomes[0]
	r.outcomes = r.outcomes[1:]
	if o.err != nil {
		return nil, o.err
	}
	return &http.Response{
		StatusCode: o.code,
		Status:     http.StatusText(o.code),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       ioutil.NopCloser(bytes.NewBufferString(o.body)),
	}, nil
}

// dnsErr satisfies isBreakerSuccess(false): a non-SDK error has no RequestId
// so it counts as a breaker failure and triggers the next candidate.
type dnsErr struct{ msg string }

func (e dnsErr) Error() string { return e.msg }

func newCvmRequest() *requestWithClientToken {
	return newTestRequest()
}

func runFailoverClient(rt http.RoundTripper) *Client {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)
	return c
}

// 15: First request DNS miss on .com → succeeds on .cn
func Test_failover_dnsMissOnComSucceedsOnCn(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	rt.programOk()
	c := runFailoverClient(rt)

	req := newCvmRequest()
	resp := tchttp.NewCommonResponse()
	if err := c.Send(req, resp); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 2 {
		t.Fatalf("expected 2 attempts, got %d (%v)", len(rt.hostsSeen), rt.hostsSeen)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com" {
		t.Fatalf("attempt[0] host = %s, want cvm.tencentcloudapi.com", rt.hostsSeen[0])
	}
	if rt.hostsSeen[1] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("attempt[1] host = %s, want cvm.tencentcloudapi.cn", rt.hostsSeen[1])
	}
	// User must see the host they configured (origin) after the call returns.
	if got := req.GetDomain(); got != "cvm.tencentcloudapi.com" {
		t.Fatalf("after failover, request domain = %s, want origin restored", got)
	}
}

// 15b: Cyclic rotation — .cn origin must roll forward to .com.cn first, then .com.
func Test_failover_cyclicRotationFromCn(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	rt.programOk()

	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.cn"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 2 {
		t.Fatalf("expected 2 attempts, got %d (%v)", len(rt.hostsSeen), rt.hostsSeen)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("attempt[0] host = %s, want cvm.tencentcloudapi.cn", rt.hostsSeen[0])
	}
	// Cyclic: .cn (idx 1) → .com.cn (idx 2) → .com (idx 0).
	if rt.hostsSeen[1] != "cvm.tencentcloudapi.com.cn" {
		t.Fatalf("attempt[1] host = %s, want cvm.tencentcloudapi.com.cn (cyclic rotation)",
			rt.hostsSeen[1])
	}
}

// 15c: Cyclic rotation — .com.cn origin → .com first, then .cn.
func Test_failover_cyclicRotationFromComCn(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	rt.programOk()

	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com.cn"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com.cn" {
		t.Fatalf("attempt[0] host = %s, want cvm.tencentcloudapi.com.cn", rt.hostsSeen[0])
	}
	// Cyclic: .com.cn (idx 2) → .com (idx 0) → .cn (idx 1).
	if rt.hostsSeen[1] != "cvm.tencentcloudapi.com" {
		t.Fatalf("attempt[1] host = %s, want cvm.tencentcloudapi.com (cyclic rotation)",
			rt.hostsSeen[1])
	}
}

// 16: All 3 TLDs error → aggregated error
func Test_failover_allTldsFail_aggregatesError(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns 1"})
	rt.programErr(dnsErr{"dns 2"})
	rt.programErr(dnsErr{"dns 3"})
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected aggregated error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "cvm.tencentcloudapi.com") ||
		!strings.Contains(msg, "cvm.tencentcloudapi.cn") ||
		!strings.Contains(msg, "cvm.tencentcloudapi.com.cn") {
		t.Fatalf("aggregated error must mention all candidates, got: %v", msg)
	}
	if len(rt.hostsSeen) != 3 {
		t.Fatalf("expected 3 attempts, got %d", len(rt.hostsSeen))
	}
}

// 17: Business error AuthFailure with RequestId → no failover, single attempt
func Test_failover_businessErrorPropagatesImmediately(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(200, `{"Response":{"Error":{"Code":"AuthFailure.SignatureExpire","Message":"expired"},"RequestId":"req-123"}}`)
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected SDK error, got nil")
	}
	if !strings.Contains(err.Error(), "AuthFailure") {
		t.Fatalf("expected AuthFailure to propagate, got %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("business error must not trigger failover, attempts = %d", len(rt.hostsSeen))
	}
}

// 18: Non-200 response on .com → failover to .cn → success
func Test_failover_non200TriggersFailover(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(503, `{}`)
	rt.programOk()
	c := runFailoverClient(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 2 {
		t.Fatalf("expected 2 attempts, got %d", len(rt.hostsSeen))
	}
	if rt.hostsSeen[1] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("expected failover to cn, got %s", rt.hostsSeen[1])
	}
}

// Invalid JSON 200 response → failover (parser fails with empty RequestId)
func Test_failover_invalidJsonTriggersFailover(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(200, `<html>blocked</html>`)
	rt.programOk()
	c := runFailoverClient(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 2 {
		t.Fatalf("expected 2 attempts, got %d", len(rt.hostsSeen))
	}
}

// 19: Subsequent request after .com breaker tripped → only .cn hit
func Test_failover_breakerOpenOnComShortCircuits(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &scriptedRT{}
	c.WithHttpTransport(rt)

	// Pre-trip the .com breaker by direct manipulation.
	state := c.fo.stateFor("cvm.tencentcloudapi.com")
	tripBreakerForTest(state.breakerFor("cvm.tencentcloudapi.com"))

	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on .cn, got %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("breaker open on .com should short-circuit, got attempts=%d (%v)",
			len(rt.hostsSeen), rt.hostsSeen)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("expected first reachable host = .cn, got %s", rt.hostsSeen[0])
	}
}

// 20: Credential rotation between attempts: second request uses new credentials.
func Test_failover_credentialRotationBetweenAttempts(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	rt.programOk()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDOLD", "SKOLD"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	// We can't synchronously interrupt mid-Send; instead make a first call
	// that exercises both attempts with the old creds, swap, then make a
	// second call and verify the new creds are used. The signature path
	// reads c.credential per call so this proves dynamic re-reading.
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("first call: %v", err)
	}
	auth1 := rt.authsSeen[0]
	if !strings.Contains(auth1, "AKIDOLD") {
		t.Fatalf("first attempt auth = %s, want AKIDOLD", auth1)
	}
	auth2 := rt.authsSeen[1]
	if !strings.Contains(auth2, "AKIDOLD") {
		t.Fatalf("resign auth = %s, want still AKIDOLD (credential not yet rotated)", auth2)
	}

	// Rotate creds.
	c.WithCredential(NewCredential("AKIDNEW", "SKNEW"))

	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("second call: %v", err)
	}
	auth3 := rt.authsSeen[2]
	if !strings.Contains(auth3, "AKIDNEW") {
		t.Fatalf("post-rotation auth = %s, want AKIDNEW", auth3)
	}
}

// 21: backupEndpoint mode: region-pinned origin failed → falls over to <service>.<backupEP>
func Test_failover_backupEndpointMode(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	rt.programOk()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-shanghai.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	// Force origin to be a region-pinned host (would otherwise be ineligible
	// for TLD rotation, but backupEndpoint overrides that).
	cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 2 {
		t.Fatalf("expected 2 attempts, got %d", len(rt.hostsSeen))
	}
	if rt.hostsSeen[0] != "cvm.ap-guangzhou.tencentcloudapi.com" {
		t.Fatalf("attempt[0] = %s, want region-pinned origin", rt.hostsSeen[0])
	}
	if rt.hostsSeen[1] != "cvm.ap-shanghai.tencentcloudapi.com" {
		t.Fatalf("attempt[1] = %s, want service.backupEndpoint", rt.hostsSeen[1])
	}
}

// Sanity: the happy path (200 + valid JSON) makes only one attempt.
func Test_failover_happyPathOneAttempt(t *testing.T) {
	rt := &scriptedRT{}
	rt.programOk()
	c := runFailoverClient(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("happy path made %d attempts, want 1", len(rt.hostsSeen))
	}
}

// Pass-through: when DisableRegionBreaker = true the failover layer is bypassed.
func Test_failover_disabled_goesThrough(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = true
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err == nil {
		t.Fatal("expected error to propagate without failover")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt with failover disabled, got %d", len(rt.hostsSeen))
	}
}

// Drive 5 sustained transport failures on .com, with .cn always succeeding,
// and confirm the next call short-circuits .com (its breaker is now Open).
func Test_failover_breakerOpensAfterSustainedFailure(t *testing.T) {
	rt := &scriptedRT{}
	c := runFailoverClient(rt)

	for i := 0; i < 5; i++ {
		rt.programErr(dnsErr{"dns miss"})
		rt.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("iteration %d: %v", i, err)
		}
	}
	if len(rt.hostsSeen) != 10 {
		t.Fatalf("expected 10 attempts after 5 cycles, got %d", len(rt.hostsSeen))
	}

	// 6th call: .com breaker should now be Open. Only .cn is hit.
	rt.hostsSeen = nil
	rt.authsSeen = nil
	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("6th call: %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("breaker open should short-circuit .com, got %d attempts", len(rt.hostsSeen))
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("expected first reachable = .cn, got %s", rt.hostsSeen[0])
	}
}

// =====================================================================
// Helpers
// =====================================================================

// tripBreakerForTest drives the breaker's internal counters until it is Open.
// Using direct state manipulation is faster than running 6 real loops.
func tripBreakerForTest(b *circuitBreaker) {
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(time.Hour)
	b.mu.Unlock()
}

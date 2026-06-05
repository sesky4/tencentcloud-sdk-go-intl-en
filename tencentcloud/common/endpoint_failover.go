package common

import (
	"fmt"
	"strings"
	"sync"
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

// breakerTimeoutMs is how long a host stays in the Open state before a
// HalfOpen probe is permitted. Mirrors Java's BREAKER_TIMEOUT_MS.
const breakerTimeoutMs = 60 * time.Second

// tldFamilies enumerates the three Tencent Cloud TLD families recognised by
// TLD-rotation mode. Region-pinned hosts skip rotation — failing them over
// silently changes the resolved region.
var tldFamilies = [3][3]string{
	{"tencentcloudapi.com", "tencentcloudapi.cn", "tencentcloudapi.com.cn"},
	{"ai.tencentcloudapi.com", "ai.tencentcloudapi.cn", "ai.tencentcloudapi.com.cn"},
	{"internal.tencentcloudapi.com", "internal.tencentcloudapi.cn", "internal.tencentcloudapi.com.cn"},
}

var (
	familyMarkers  = []string{"ai", "internal"} // index 0 → familyIdx 1
	baseTLDs       = tldFamilies[0]
	regionPrefixes = []string{"ap-", "na-", "eu-", "sa-", "af-", "me-"}
)

// tldMatch describes how a host fits the Tencent Cloud TLD scheme.
type tldMatch struct {
	familyIdx     int    // 0=plain, 1=ai, 2=internal
	tldIdx        int    // 0=.com, 1=.cn, 2=.com.cn
	hasRegion     bool
	servicePrefix string // host's service-name labels joined by '.'
}

// failoverState holds per-origin breakers keyed by candidate host. State is
// per-origin so that tests and operators can isolate the lifecycle of one
// host's breaker without affecting others.
type failoverState struct {
	mu       sync.Mutex
	timeout  time.Duration
	breakers map[string]*circuitBreaker
}

func (s *failoverState) breakerFor(host string) *circuitBreaker {
	s.mu.Lock()
	defer s.mu.Unlock()
	if b, ok := s.breakers[host]; ok {
		return b
	}
	b := newCircuitBreaker(breakerSetting{
		maxFailNum:          defaultMaxFailNum,
		maxFailPercentage:   defaultMaxFailPercentage,
		windowInterval:      defaultWindowLength,
		timeout:             s.timeout,
		halfOpenMaxRequests: defaultHalfOpenMaxRequests,
	})
	s.breakers[host] = b
	return b
}

// endpointFailover is per-Client; per-origin breaker state lives inside it.
type endpointFailover struct {
	mu        sync.Mutex
	perOrigin map[string]*failoverState
}

func newEndpointFailover() *endpointFailover {
	return &endpointFailover{perOrigin: map[string]*failoverState{}}
}

func (f *endpointFailover) stateFor(originHost string) *failoverState {
	f.mu.Lock()
	defer f.mu.Unlock()
	if s, ok := f.perOrigin[originHost]; ok {
		return s
	}
	s := &failoverState{
		timeout:  breakerTimeoutMs,
		breakers: map[string]*circuitBreaker{},
	}
	f.perOrigin[originHost] = s
	return s
}

// ----------------------------------------------------------------------
// Host classification
// ----------------------------------------------------------------------

// tldMatchOf recognises host = "<prefix>.<base TLD>" where prefix mixes
// service labels with optional family markers and a regional label. Returns
// nil if no base TLD matches.
func tldMatchOf(host string) *tldMatch {
	if host == "" {
		return nil
	}
	baseIdx := matchBaseTld(host)
	if baseIdx < 0 {
		return nil
	}
	prefix := host[:len(host)-len(baseTLDs[baseIdx])-1]

	familyIdx := 0
	hasRegion := false
	var serviceParts []string
	for _, label := range strings.Split(prefix, ".") {
		if looksLikeRegionLabel(label) {
			hasRegion = true
			continue
		}
		if mi := familyMarkerIdx(label); mi > 0 {
			familyIdx = mi
			continue
		}
		serviceParts = append(serviceParts, label)
	}
	return &tldMatch{
		familyIdx:     familyIdx,
		tldIdx:        baseIdx,
		hasRegion:     hasRegion,
		servicePrefix: strings.Join(serviceParts, "."),
	}
}

// matchBaseTld returns the index of the longest baseTLDs entry suffixing
// host, or -1 if none match.
func matchBaseTld(host string) int {
	best := -1
	bestLen := -1
	for i, b := range baseTLDs {
		suffix := "." + b
		if !strings.HasSuffix(host, suffix) || len(suffix) <= bestLen {
			continue
		}
		prefix := host[:len(host)-len(suffix)]
		if prefix == "" || strings.HasPrefix(prefix, ".") || strings.HasSuffix(prefix, ".") {
			continue
		}
		best = i
		bestLen = len(suffix)
	}
	return best
}

func familyMarkerIdx(label string) int {
	for i, m := range familyMarkers {
		if m == label {
			return i + 1
		}
	}
	return 0
}

func looksLikeRegionLabel(label string) bool {
	for _, p := range regionPrefixes {
		if strings.HasPrefix(label, p) {
			return true
		}
	}
	return false
}

func serviceOf(host string) string {
	if i := strings.IndexByte(host, '.'); i >= 0 {
		return host[:i]
	}
	return host
}

// ----------------------------------------------------------------------
// Candidate selection
// ----------------------------------------------------------------------

type candidate struct {
	host    string
	breaker *circuitBreaker
}

// candidatesFor returns the ordered candidate list for originHost, or
// nil if the host is not eligible for failover (caller should pass
// through to plain sendWithSignature).
func candidatesFor(c *Client, originHost string) []candidate {
	state := c.fo.stateFor(originHost)

	if backupEP := backupEndpointOf(c.profile); backupEP != "" {
		backupHost := serviceOf(originHost) + "." + backupEP
		return []candidate{
			{host: originHost, breaker: state.breakerFor(originHost)},
			{host: backupHost, breaker: state.breakerFor(backupHost)},
		}
	}

	m := tldMatchOf(originHost)
	if m == nil || m.hasRegion {
		return nil
	}

	// Cyclic rotation starting at origin: origin, (origin+1) % n, …
	// Order is independent of which TLD the user configured —
	// .cn → [.cn, .com.cn, .com], not [.cn, .com, .com.cn].
	n := len(baseTLDs)
	cs := make([]candidate, 0, n)
	for i := 0; i < n; i++ {
		t := (m.tldIdx + i) % n
		host := m.servicePrefix + "." + tldFamilies[m.familyIdx][t]
		cs = append(cs, candidate{host: host, breaker: state.breakerFor(host)})
	}
	return cs
}

func backupEndpointOf(p *profile.ClientProfile) string {
	if p == nil {
		return ""
	}
	if p.BackupEndpoint != "" {
		return p.BackupEndpoint
	}
	return p.BackupEndPoint
}

// ----------------------------------------------------------------------
// Execution
// ----------------------------------------------------------------------

// sendWithEndpointFailover walks the candidate list, driving each
// candidate's breaker, until one succeeds or all are exhausted. The origin
// host is restored before returning so callers see what they configured.
func (c *Client) sendWithEndpointFailover(request tchttp.Request, response tchttp.Response) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("endpoint failover panic: %v", e)
		}
	}()

	originHost := request.GetDomain()
	candidates := candidatesFor(c, originHost)
	if candidates == nil {
		return c.sendWithSignature(request, response)
	}
	defer request.SetDomain(originHost)

	var attemptErrs []string
	for _, cand := range candidates {
		gen, openErr := cand.breaker.beforeRequest()
		if openErr != nil {
			attemptErrs = append(attemptErrs, fmt.Sprintf("%s: circuit breaker open", cand.host))
			continue
		}
		request.SetDomain(cand.host)
		err = c.sendWithSignature(request, response)
		success := isBreakerSuccess(err)
		cand.breaker.afterRequest(gen, success)
		if success {
			return err
		}
		attemptErrs = append(attemptErrs,
			fmt.Sprintf("%s: %v", cand.host, err))
	}

	if err == nil {
		// Every candidate was breaker-skipped; no real attempt produced
		// an error. Surface a synthetic one so the caller can see why.
		err = fmt.Errorf("endpoint failover: all candidates skipped for %s", originHost)
	}
	if len(attemptErrs) > 1 {
		err = fmt.Errorf("endpoint failover failed for %s: %s",
			originHost, strings.Join(attemptErrs, "; "))
	}
	return err
}

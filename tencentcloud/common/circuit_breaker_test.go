package common

import (
	"errors"
	"testing"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
)

// Test_counter_onFailure_incrementsConsecutiveFailures verifies that failures
// update the consecutive-failure counter used by readyToOpen.
func Test_counter_onFailure_incrementsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	for i := 0; i < 6; i++ {
		c.onFailure()
	}
	if c.consecutiveFailures != 6 {
		t.Fatalf("after 6 onFailure calls, consecutiveFailures = %d, want 6", c.consecutiveFailures)
	}
	if c.failures != 6 || c.all != 6 {
		t.Fatalf("after 6 onFailure calls, failures=%d all=%d, want 6/6", c.failures, c.all)
	}
}

// Test_counter_onSuccess_resetsConsecutiveFailures verifies that a single success
// resets consecutiveFailures back to zero (the well-known burst recovery contract).
func Test_counter_onSuccess_resetsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	c.onFailure()
	c.onFailure()
	c.onSuccess()
	if c.consecutiveFailures != 0 {
		t.Fatalf("after onSuccess, consecutiveFailures = %d, want 0", c.consecutiveFailures)
	}
}

// Test_counter_clear_resetsConsecutiveFailures verifies that generation rollovers
// do not carry over a stale consecutive-failure count.
func Test_counter_clear_resetsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	c.onFailure()
	c.onFailure()
	c.clear()
	if c.consecutiveFailures != 0 {
		t.Fatalf("after clear, consecutiveFailures = %d, want 0", c.consecutiveFailures)
	}
	if c.failures != 0 || c.all != 0 || c.consecutiveSuccesses != 0 {
		t.Fatalf("after clear, failures=%d all=%d consecutiveSuccesses=%d, want all 0",
			c.failures, c.all, c.consecutiveSuccesses)
	}
}

// Test_circuitBreaker_tripsOnConsecutiveFailures isolates the consecutive-failures
// trip branch (readyToOpen second clause: consecutiveFailures > 5).
//
// To isolate from the rate-based branch we need failures/all < 75%. We use a custom
// breakerSetting with maxFailNum=1000 (effectively disabling the rate branch) so that
// only the consecutive-failures path can trip. We then send 6 failures in a row and
// assert the breaker has opened.
func Test_circuitBreaker_tripsOnSixConsecutiveFailures(t *testing.T) {
	cb := newCircuitBreaker(breakerSetting{
		maxFailNum:        1000, // suppress the rate-based branch
		maxFailPercentage: 75,
		windowInterval:    1 * time.Hour,
		timeout:           60 * time.Second,
	})
	for i := 0; i < 6; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: beforeRequest returned err=%v before breaker should open", i, err)
		}
		cb.afterRequest(gen, false)
	}
	// After 6 consecutive failures the breaker must be open via the consecutive-failures branch.
	_, err := cb.beforeRequest()
	if err != errOpenState {
		t.Fatalf("after 6 consecutive failures, breaker should be open, got err=%v", err)
	}
}

// Test_isBreakerSuccess_nilErrIsSuccess verifies that a successful API call
// counts as a success to the circuit breaker.
func Test_isBreakerSuccess_nilErrIsSuccess(t *testing.T) {
	if !isBreakerSuccess(nil) {
		t.Fatal("isBreakerSuccess(nil) = false, want true (nil error is a successful call)")
	}
}

// Test_isBreakerSuccess_sdkErrorWithRequestIdIsSuccess: a structured server error
// with a RequestId still indicates the region answered, so the breaker should treat it
// as success unless it is an InternalError (which marks the region itself unhealthy).
func Test_isBreakerSuccess_sdkErrorWithRequestIdIsSuccess(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("AuthFailure.SignatureExpire", "expired", "req-123")
	if !isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(AuthFailure with RequestId) = false, want true")
	}
}

// Test_isBreakerSuccess_internalErrorIsFailure: InternalError indicates region-side
// fault; treat as a failure for breaker accounting.
func Test_isBreakerSuccess_internalErrorIsFailure(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("InternalError", "boom", "req-123")
	if isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(InternalError) = true, want false")
	}
}

// Test_isBreakerSuccess_sdkErrorWithoutRequestIdIsFailure: missing RequestId means
// the response never reached the user from the region (locally-fabricated error,
// network failure mid-flight, etc.), so the breaker should consider it a failure.
func Test_isBreakerSuccess_sdkErrorWithoutRequestIdIsFailure(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("ClientError.NetworkError", "boom", "")
	if isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(SDKError without RequestId) = true, want false")
	}
}

// Test_isBreakerSuccess_nonSDKErrorIsFailure: arbitrary network/transport error must
// be a failure.
func Test_isBreakerSuccess_nonSDKErrorIsFailure(t *testing.T) {
	if isBreakerSuccess(errors.New("connection refused")) {
		t.Fatal("isBreakerSuccess(non-SDK error) = true, want false")
	}
}

// Test_circuitBreaker_doesNotOpenOnSuccesses verifies that successful calls do
// not open the breaker and trigger wrong-region routing.
func Test_circuitBreaker_doesNotOpenOnSuccesses(t *testing.T) {
	cb := newCircuitBreaker(breakerSetting{
		maxFailNum:          defaultMaxFailNum,
		maxFailPercentage:   defaultMaxFailPercentage,
		windowInterval:      defaultWindowLength,
		timeout:             defaultTimeout,
		halfOpenMaxRequests: defaultHalfOpenMaxRequests,
	})
	for i := 0; i < 10; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: breaker opened spuriously, err=%v", i, err)
		}
		// Simulate a successful call: nil error -> isBreakerSuccess returns true.
		cb.afterRequest(gen, isBreakerSuccess(nil))
	}
	// Confirm state is still closed.
	cb.mu.Lock()
	st := cb.state
	cb.mu.Unlock()
	if st != StateClosed {
		t.Fatalf("after 10 successes, breaker state = %v, want StateClosed", st)
	}
}

// guards against the future regression of the dead-code branch by waiting through
// a generation rollover and confirming consecutiveFailures has been cleared. Avoids
// a real time.Sleep by configuring an extremely short windowInterval.
func Test_circuitBreaker_generationRollover_clearsConsecutiveFailures(t *testing.T) {
	cb := newCircuitBreaker(breakerSetting{
		maxFailNum:        5,
		maxFailPercentage: 75,
		windowInterval:    1 * time.Millisecond,
		timeout:           60 * time.Second,
	})
	// Record a few failures (not enough to trip).
	for i := 0; i < 3; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: unexpected err=%v", i, err)
		}
		cb.afterRequest(gen, false)
	}
	// Wait past windowInterval so the next beforeRequest rolls a new generation.
	time.Sleep(5 * time.Millisecond)
	gen, err := cb.beforeRequest()
	if err != nil {
		t.Fatalf("after rollover, unexpected err=%v", err)
	}
	_ = gen
	cb.mu.Lock()
	cf := cb.counter.consecutiveFailures
	all := cb.counter.all
	failures := cb.counter.failures
	cb.mu.Unlock()
	if cf != 0 || all != 0 || failures != 0 {
		t.Fatalf("after generation rollover counter not cleared: consecutiveFailures=%d failures=%d all=%d",
			cf, failures, all)
	}
}

func Test_circuitBreaker_halfOpenLimitsProbeRequests(t *testing.T) {
	cb := newCircuitBreaker(breakerSetting{
		maxFailNum:          5,
		maxFailPercentage:   75,
		windowInterval:      time.Hour,
		timeout:             time.Second,
		halfOpenMaxRequests: 2,
	})

	cb.state = StateOpen
	cb.expiry = time.Now().Add(-time.Second)

	gen1, err := cb.beforeRequest()
	if err != nil {
		t.Fatalf("first half-open probe was rejected: %v", err)
	}
	gen2, err := cb.beforeRequest()
	if err != nil {
		t.Fatalf("second half-open probe was rejected: %v", err)
	}
	rejectedGen, err := cb.beforeRequest()
	if err != errOpenState {
		t.Fatalf("third half-open probe should be rejected, got %v", err)
	}
	cb.afterRequest(rejectedGen, true)
	cb.mu.Lock()
	inFlightAfterRejectedBackup := cb.halfOpenInFlight
	cb.mu.Unlock()
	if inFlightAfterRejectedBackup != 2 {
		t.Fatalf("rejected half-open backup request changed in-flight probes to %d, want 2", inFlightAfterRejectedBackup)
	}

	cb.afterRequest(gen1, true)
	cb.mu.Lock()
	stateAfterOneSuccess := cb.state
	cb.mu.Unlock()
	if stateAfterOneSuccess != StateHalfOpen {
		t.Fatalf("after one successful probe, state=%v, want StateHalfOpen", stateAfterOneSuccess)
	}

	cb.afterRequest(gen2, true)
	cb.mu.Lock()
	stateAfterTwoSuccesses := cb.state
	cb.mu.Unlock()
	if stateAfterTwoSuccesses != StateClosed {
		t.Fatalf("after two successful probes, state=%v, want StateClosed", stateAfterTwoSuccesses)
	}
}

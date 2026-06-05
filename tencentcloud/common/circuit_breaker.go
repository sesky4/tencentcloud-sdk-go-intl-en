// a simple golang breaker according to https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10)

package common

import (
	"errors"
	"sync"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
)

const (
	defaultMaxFailNum          = 5
	defaultMaxFailPercentage   = 75
	defaultWindowLength        = 1 * 60 * time.Second
	defaultTimeout             = 60 * time.Second
	defaultHalfOpenMaxRequests = 1
)

const ignoredGeneration = ^uint64(0)

var (
	// ErrOpenState is returned when the CB state is open
	errOpenState = errors.New("circuit breaker is open")
)

// counter is protected by circuitBreaker.mu.
type counter struct {
	failures             int
	all                  int
	consecutiveSuccesses int
	consecutiveFailures  int
}

func newRegionCounter() counter {
	return counter{
		failures:             0,
		all:                  0,
		consecutiveSuccesses: 0,
		consecutiveFailures:  0,
	}
}

func (c *counter) onSuccess() {
	c.all++
	c.consecutiveSuccesses++
	c.consecutiveFailures = 0
}

func (c *counter) onFailure() {
	c.all++
	c.failures++
	c.consecutiveSuccesses = 0
	c.consecutiveFailures++
}

func (c *counter) clear() {
	c.all = 0
	c.failures = 0
	c.consecutiveSuccesses = 0
	c.consecutiveFailures = 0
}

// State is a type that represents a state of CircuitBreaker.
type state int

// These constants are states of CircuitBreaker.
const (
	StateClosed state = iota
	StateHalfOpen
	StateOpen
)

type breakerSetting struct {
	// max fail nums
	// the default is 5
	maxFailNum int
	// max fail percentage
	// the default is 75/100
	maxFailPercentage int
	// windowInterval decides when to reset counter if the state is StateClosed
	// the default is 60s
	windowInterval time.Duration
	// timeout decides when to turn StateOpen to StateHalfOpen
	// the default is 60s
	timeout time.Duration
	// halfOpenMaxRequests limits probe requests allowed in StateHalfOpen.
	halfOpenMaxRequests int
}

type circuitBreaker struct {
	// settings
	breakerSetting
	// read and write lock
	mu sync.Mutex
	// the breaker's state: closed, open, half-open
	state state
	// expiry time determines whether to enter the next generation
	// if in StateClosed, it will be now + windowInterval
	// if in StateOpen, it will be now + timeout
	// if in StateHalfOpen. it will be zero
	expiry time.Time
	// generation decide whether add the afterRequest's request to counter
	generation uint64
	// counter
	counter counter
	// halfOpenInFlight tracks probe requests admitted in StateHalfOpen.
	halfOpenInFlight int
}

func newCircuitBreaker(set breakerSetting) (re *circuitBreaker) {
	if set.halfOpenMaxRequests <= 0 {
		set.halfOpenMaxRequests = defaultHalfOpenMaxRequests
	}
	re = new(circuitBreaker)
	re.breakerSetting = set
	return
}

// currentState return the current state.
//
//	if in StateClosed and now is over expiry time, it will turn to a new generation.
//	if in StateOpen and now is over expiry time, it will turn to StateHalfOpen
func (s *circuitBreaker) currentState(now time.Time) (state, uint64) {
	switch s.state {
	case StateClosed:
		if s.expiry.Before(now) {
			s.toNewGeneration(now)
		}
	case StateOpen:
		if s.expiry.Before(now) {
			s.setState(StateHalfOpen, now)
		}
	}
	return s.state, s.generation
}

// setState set the circuitBreaker's state to newState
// and turn to new generation
func (s *circuitBreaker) setState(newState state, now time.Time) {
	if s.state == newState {
		return
	}
	s.state = newState
	s.toNewGeneration(now)
}

// toNewGeneration will increase the generation and clear the counter.
// it also will reset the expiry
func (s *circuitBreaker) toNewGeneration(now time.Time) {
	s.generation++
	s.counter.clear()
	s.halfOpenInFlight = 0
	var zero time.Time
	switch s.state {
	case StateClosed:
		s.expiry = now.Add(s.windowInterval)
	case StateOpen:
		s.expiry = now.Add(s.timeout)
	default: // StateHalfOpen
		s.expiry = zero
	}
}

// beforeRequest return the current generation; if the breaker is in StateOpen, it will also return an errOpenState
func (s *circuitBreaker) beforeRequest() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	state, generation := s.currentState(now)
	//log.Println(s.counter)
	switch state {
	case StateOpen:
		return ignoredGeneration, errOpenState
	case StateHalfOpen:
		if s.counter.all+s.halfOpenInFlight >= s.halfOpenMaxRequests {
			return ignoredGeneration, errOpenState
		}
		s.halfOpenInFlight++
	}
	return generation, nil
}

func (s *circuitBreaker) afterRequest(before uint64, success bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	state, generation := s.currentState(now)
	// the breaker has entered the next generation, the current results abandon.
	if generation != before {
		return
	}
	if state == StateHalfOpen && s.halfOpenInFlight > 0 {
		s.halfOpenInFlight--
	}
	if success {
		s.onSuccess(state, now)
	} else {
		s.onFailure(state, now)
	}
}

func (s *circuitBreaker) onSuccess(state state, now time.Time) {
	switch state {
	case StateClosed:
		s.counter.onSuccess()
	case StateHalfOpen:
		s.counter.onSuccess()
		// The conditions for closing breaker are met
		if s.counter.all-s.counter.failures >= s.halfOpenMaxRequests {
			s.setState(StateClosed, now)
		}
	}
}

func (s *circuitBreaker) readyToOpen(c counter) bool {
	failPre := float64(c.failures) / float64(c.all)
	return (c.failures >= s.maxFailNum && failPre >= float64(s.maxFailPercentage)/100.0) ||
		c.consecutiveFailures > 5
}

func (s *circuitBreaker) onFailure(state state, now time.Time) {
	switch state {
	case StateClosed:
		s.counter.onFailure()
		if f := s.readyToOpen(s.counter); f {
			s.setState(StateOpen, now)
		}
	case StateHalfOpen:
		s.setState(StateOpen, now)
	}
}

// isBreakerSuccess decides whether a sendWithSignature result counts as a
// success for the regional circuit breaker.
//
// The breaker tracks region health, not per-call business outcome. So:
//   - nil err: the request round-tripped and parsed cleanly → region healthy.
//   - *TencentCloudSDKError with a RequestId: the region answered with a
//     structured error (e.g. AuthFailure). The user's call failed but the
//     region is up → success for breaker accounting. The single exception is
//     "InternalError", which the API uses to signal a region-side fault.
//   - Anything else (transport errors, locally-fabricated errors with no
//     RequestId): the region did not answer → failure.
func isBreakerSuccess(err error) bool {
	if err == nil {
		return true
	}
	e, ok := err.(*tcerr.TencentCloudSDKError)
	if !ok {
		return false
	}
	if e.GetRequestId() == "" {
		return false
	}
	if e.GetCode() == "InternalError" {
		return false
	}
	return true
}

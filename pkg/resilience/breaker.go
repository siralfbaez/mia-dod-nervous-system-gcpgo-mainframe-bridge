package resilience

import (
	"errors"
	"sync"
	"time"
)

var ErrServiceUnavailable = errors.New("circuit breaker is open: service unavailable")

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

// CircuitBreaker protects the Mainframe bridge from cascading failures
type CircuitBreaker struct {
	mutex           sync.Mutex
	state           State
	failureCount    int
	failureThreshold int
	resetTimeout    time.Duration
	lastFailureTime time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: threshold,
		resetTimeout:     timeout,
	}
}

// Execute wraps a call to a legacy system (e.g., C++ Mainframe service)
func (cb *CircuitBreaker) Execute(request func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	
	// Check if we should transition from Open to Half-Open
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			cb.mutex.Unlock()
			return nil, ErrServiceUnavailable
		}
	}
	cb.mutex.Unlock()

	// Perform the actual work
	result, err := request()

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.onFailure()
		return nil, err
	}

	cb.onSuccess()
	return result, nil
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	if cb.failureCount >= cb.failureThreshold {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) onSuccess() {
	cb.failureCount = 0
	cb.state = StateClosed
}

package resilience

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	threshold := 3
	timeout := 100 * time.Millisecond
	cb := NewCircuitBreaker(threshold, timeout)

	// Simulated failing function
	failFunc := func() (interface{}, error) {
		return nil, errors.New("mainframe timeout")
	}

	// 1. Trigger failures until threshold reached
	for i := 0; i < threshold; i++ {
		_, err := cb.Execute(failFunc)
		if err == nil {
			t.Errorf("Expected error at attempt %d, got nil", i+1)
		}
	}

	// 2. The circuit should now be OPEN
	_, err := cb.Execute(failFunc)
	if err != ErrServiceUnavailable {
		t.Error("Expected ErrServiceUnavailable (Open Circuit), got:", err)
	}

	// 3. Wait for reset timeout
	time.Sleep(timeout + 10*time.Millisecond)

	// 4. Circuit should be Half-Open. Let's simulate a success.
	successFunc := func() (interface{}, error) {
		return "success", nil
	}

	result, err := cb.Execute(successFunc)
	if err != nil || result != "success" {
		t.Errorf("Expected success on half-open reset, got: %v", err)
	}

	// 5. Circuit should be CLOSED again
	if cb.state != StateClosed {
		t.Error("Circuit should have transitioned back to StateClosed")
	}
}

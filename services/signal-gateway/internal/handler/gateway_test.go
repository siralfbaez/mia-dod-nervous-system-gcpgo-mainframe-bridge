package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/pkg/resilience"
)

func TestProcessMainframeSignal_BreakerOpen(t *testing.T) {
	// 1. Initialize a breaker and manually trip it
	breaker := resilience.NewCircuitBreaker(1, 10*time.Second)

	// Force failure to open the circuit
	breaker.Execute(func() (interface{}, error) {
		return nil, resilience.ErrServiceUnavailable
	})

	h := &SignalHandler{Breaker: breaker}

	// 2. Create a request
	req := httptest.NewRequest("POST", "/v1/signal", strings.NewReader(`{"test":"data"}`))
	rr := httptest.NewRecorder()

	// 3. Execute handler
	h.ProcessMainframeSignal(rr, req)

	// 4. Check for 503 Service Unavailable
	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("Expected 503 Service Unavailable when breaker is open, got %d", rr.Code)
	}
}
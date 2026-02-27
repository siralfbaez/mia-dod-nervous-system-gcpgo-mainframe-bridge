package handler

import (
	"encoding/json"
	"net/http"
	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/pkg/resilience"
)

type SignalHandler struct {
	Breaker *resilience.CircuitBreaker
}

func (h *SignalHandler) ProcessMainframeSignal(w http.ResponseWriter, r *http.Request) {
	// Wrap the processing in the Circuit Breaker
	_, err := h.Breaker.Execute(func() (interface{}, error) {
		// 1. Decode legacy payload
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return nil, err
		}

		// 2. Logic to forward to Translation Engine would go here
		// For now, we simulate success
		return "Signal Buffered", nil
	})

	if err != nil {
		if err == resilience.ErrServiceUnavailable {
			http.Error(w, "Mainframe Bridge Saturated - Try Again Later", http.StatusServiceUnavailable)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}

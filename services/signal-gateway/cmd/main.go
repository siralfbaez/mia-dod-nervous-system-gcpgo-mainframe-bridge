package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/pkg/resilience"
	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/services/signal-gateway/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the Circuit Breaker for the Gateway
	breaker := resilience.NewCircuitBreaker(5, 30*time.Second)

	h := &handler.SignalHandler{
		Breaker: breaker,
	}

	http.HandleFunc("/v1/signal", h.ProcessMainframeSignal)

	log.Printf("📡 Signal Gateway Pulse active on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
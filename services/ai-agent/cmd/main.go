package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/pkg/vertexai"
	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/services/ai-agent/internal/logic"
)

func main() {
	// 1. Setup Context and Configuration
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	projectID := os.Getenv("GCP_PROJECT_ID")
	region    := os.Getenv("GCP_REGION") // e.g., us-central1
	modelName := "gemini-1.5-pro"

	if projectID == "" {
		log.Fatal("GCP_PROJECT_ID environment variable is required")
	}

	// 2. Initialize the AI Client (The Brain's Synapses)
	aiClient, err := vertexai.NewClient(ctx, projectID, region, modelName)
	if err != nil {
		log.Fatalf("Failed to initialize Vertex AI client: %v", err)
	}

	// 3. Initialize Diagnostic Logic
	agent := &logic.DiagnosticAgent{
		AIClient: aiClient,
	}

	// 4. Simulate receiving a signal from the Mainframe Bridge
	// In a real scenario, this would come from a Pub/Sub listener
	signal := "CIRCUIT_BREAKER_OPEN: Mainframe DB2 latency > 5000ms during EBCDIC translation"
	
	log.Printf("🧠 AI Agent analyzing signal: %s", signal)

	diagnosis, err := agent.AnalyzeSystemHealth(ctx, signal)
	if err != nil {
		log.Fatalf("Diagnostic failure: %v", err)
	}

	log.Printf("✅ AI Diagnosis & Mitigation Plan: \n%s", diagnosis)
}

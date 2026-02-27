package logic

import (
	"context"
	"fmt"
	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo-mainframe-bridge/pkg/vertexai"
)

type DiagnosticAgent struct {
	AIClient *vertexai.Client
}

// AnalyzeSystemHealth takes a signal (like a circuit breaker trip) and diagnoses the root cause
func (a *DiagnosticAgent) AnalyzeSystemHealth(ctx context.Context, signal string) (string, error) {
	// Construct the prompt for the "Brain"
	prompt := fmt.Sprintf(`
		Context: You are the Diagnostic Brain for an Airline Mainframe-to-GCP Bridge.
		Signal Received: %s
		Task: Analyze if this is a Transient Network Issue, a Mainframe MIPS saturation, 
		or an Encoding mismatch (EBCDIC/UTF-8). Suggest immediate mitigation steps.
	`, signal)

	// Call Vertex AI / Gemini
	diagnosis, err := a.AIClient.GenerateResponse(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("AI brain failed to process signal: %v", err)
	}

	return diagnosis, nil
}

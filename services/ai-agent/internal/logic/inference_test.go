package logic

import (
	"context"
	"testing"
)

func TestAnalyzeSystemHealth_EmptySignal(t *testing.T) {
	agent := &DiagnosticAgent{AIClient: nil} // Mocking client
	ctx := context.Background()

	_, err := agent.AnalyzeSystemHealth(ctx, "")
	if err == nil {
		t.Error("Expected error when providing an empty signal to the AI brain")
	}
}
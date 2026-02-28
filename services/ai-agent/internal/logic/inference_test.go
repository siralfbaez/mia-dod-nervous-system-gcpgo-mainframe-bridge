package logic

import (
	"context"
	"testing"
)

func TestAnalyzeSystemHealth_EmptySignal(t *testing.T) {
	// Passing nil is now safe because of our guard clause
	agent := &DiagnosticAgent{AIClient: nil}
	ctx := context.Background()

	_, err := agent.AnalyzeSystemHealth(ctx, "")
	if err == nil {
		t.Error("Expected error when providing an empty signal or nil client, but got none")
	}
}
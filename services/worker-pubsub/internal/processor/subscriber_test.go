package processor

import (
	"context"
	"testing"
)

func TestProcessMessage_InvalidJSON(t *testing.T) {
	p := &RecordProcessor{DB: nil}
	err := p.ProcessMessage(context.Background(), []byte(`{invalid-json}`))
	if err == nil {
		t.Error("Expected error for malformed JSON payload")
	}
}
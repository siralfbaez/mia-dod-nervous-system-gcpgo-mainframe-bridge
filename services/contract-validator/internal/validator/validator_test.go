package validator

import "testing"

func TestValidatePNR(t *testing.T) {
	tests := []struct {
		pnr     string
		wantErr bool
	}{
		{"AFB123", false}, // Valid
		{"123456", false}, // Valid alphanumeric
		{"ABCDE", true},   // Too short
		{"ABCDEFG", true}, // Too long
		{"ABC!@#", true},  // Special chars
	}

	for _, tt := range tests {
		err := ValidatePNR(tt.pnr)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidatePNR(%s) error = %v, wantErr %v", tt.pnr, err, tt.wantErr)
		}
	}
}
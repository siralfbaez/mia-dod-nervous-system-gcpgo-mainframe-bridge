package encoding

import "testing"

func TestMapEBCDICToUTF8(t *testing.T) {
	// "ALF   " in EBCDIC (CCSID 37)
	// A=C1, L=D3, F=C6, Space=40
	ebcdicInput := []byte{0xC1, 0xD3, 0xC6, 0x40, 0x40, 0x40}
	expected := "ALF" // Trimmed UTF-8

	result, err := MapEBCDICToUTF8(ebcdicInput, 37)
	if err != nil {
		t.Fatalf("Failed to map EBCDIC: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestUnsupportedCCSID(t *testing.T) {
	_, err := MapEBCDICToUTF8([]byte{0x00}, 999) // Invalid CCSID
	if err == nil {
		t.Error("Expected error for unsupported CCSID, but got nil")
	}
}
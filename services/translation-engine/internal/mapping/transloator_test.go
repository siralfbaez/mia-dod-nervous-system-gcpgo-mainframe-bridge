package mapping

import "testing"

func TestRecordTranslator(t *testing.T) {
	rt := NewRecordTranslator(CCSID_US_CANADA)

	// Raw EBCDIC for "PNR123"
	raw := []byte{0xD7, 0xD5, 0xD9, 0xF1, 0xF2, 0xF3}
	expected := "PNR123"

	result, err := rt.TransformByteStream(raw)
	if err != nil {
		t.Fatalf("Translator error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
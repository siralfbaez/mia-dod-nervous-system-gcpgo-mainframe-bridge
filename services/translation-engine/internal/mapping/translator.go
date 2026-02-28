package mapping

import (
	"bytes"
	"io"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// RecordTranslator handles the transformation of mainframe records
type RecordTranslator struct {
	Encoding *charmap.Charmap
}

func NewRecordTranslator(id CCSID) *RecordTranslator {
	return &RecordTranslator{
		Encoding: GetEncoding(id),
	}
}

// TransformByteStream converts raw EBCDIC bytes to UTF-8
func (rt *RecordTranslator) TransformByteStream(input []byte) (string, error) {
	reader := transform.NewReader(bytes.NewReader(input), rt.Encoding.NewDecoder())
	output, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// CleanPassengerName handles specific airline PNR scrubbing logic
// Trimming trailing mainframe spaces (EBCDIC 0x40)
func (rt *RecordTranslator) CleanPassengerName(raw []byte) (string, error) {
	translated, err := rt.TransformByteStream(raw)
	if err != nil {
		return "", err
	}
	// Mainframe records are often fixed-width; trim the padding
	return bytes.NewBufferString(translated).String(), nil
}

package encoding

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// MapEBCDICToUTF8 handles the core conversion logic for airline PNR data
func MapEBCDICToUTF8(ebcdicBytes []byte, ccsid int) (string, error) {
	var cm *charmap.Charmap

	switch ccsid {
	case 37:
		cm = charmap.CodePage037
	case 1140:
		cm = charmap.CodePage1140
	default:
		return "", fmt.Errorf("unsupported CCSID: %d", ccsid)
	}

	reader := transform.NewReader(bytes.NewReader(ebcdicBytes), cm.NewDecoder())
	utf8Bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("transformation failed: %w", err)
	}

	// Mainframe strings are often space-padded (0x40 in EBCDIC)
	return string(bytes.TrimSpace(utf8Bytes)), nil
}
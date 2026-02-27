package mapping

import "golang.org/x/text/encoding/charmap"

// CCSID represents the IBM Coded Character Set Identifier
type CCSID int

const (
	// IBM037 is the standard EBCDIC for US/Canada
	CCSID_US_CANADA CCSID = 37
	// IBM1140 is US/Canada with the Euro symbol
	CCSID_US_EURO   CCSID = 1140
)

// GetEncoding returns the appropriate charmap for a given CCSID
func GetEncoding(id CCSID) *charmap.Charmap {
	switch id {
	case CCSID_US_EURO:
		return charmap.CodePage1140
	case CCSID_US_CANADA:
		return charmap.CodePage037
	default:
		return charmap.CodePage037 // Default to US standard
	}
}

package validator

import (
	"strings"
	"unicode"
)

// Address structure
type Address struct {
	string
}

// MarshalCSV validates the address string but does no replacement.
func (addr *Address) MarshalCSV() (string, error) {
	return strings.ToValidUTF8(addr.string, ""), nil
}

// UnmarshalCSV validates the address string from the csv file but does no
// replacement.
func (addr *Address) UnmarshalCSV(csv string) error {
	addr.string = strings.ToValidUTF8(csv, "")
	return nil
}

// Notes structure
type Notes struct {
	string
}

// MarshalCSV validates the notes string and replaces invalid characters with
// the Unicode Replacement Character.
func (notes *Notes) MarshalCSV() (string, error) {
	return strings.ToValidUTF8(notes.string, string(unicode.ReplacementChar)), nil
}

// UnmarshalCSV validates the notes string from the csv file and replaces
// invalid characters with the Unicode Replacement Character.
func (notes *Notes) UnmarshalCSV(csv string) error {
	notes.string = strings.ToValidUTF8(csv, string(unicode.ReplacementChar))
	return nil
}

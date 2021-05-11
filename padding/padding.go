package padding

import (
	"strings"
	"unicode/utf8"
)

// A ZipCode structure
type ZipCode struct {
	string
}

// MarshalCSV returns the zip code unchanged
func (zip *ZipCode) MarshalCSV() (string, error) {
	return zip.string, nil
}

// UnmarshalCSV accepts a string and unmarshals the string into a zero
// left-padded string
func (zip *ZipCode) UnmarshalCSV(csv string) error {
	maxZipCodeLength := 5
	paddingString := "0"
	currentStringLength := utf8.RuneCountInString(csv)

	if currentStringLength < maxZipCodeLength {
		zip.string = strings.Repeat(paddingString, maxZipCodeLength-currentStringLength) + csv
	} else {
		zip.string = csv
	}

	return nil
}

package capper

import "strings"

// Uppercase structure
type Uppercase struct {
	string
}

// MarshalCSV returns an uppercase string
func (fName *Uppercase) MarshalCSV() (string, error) {
	return strings.ToUpper(fName.string), nil
}

// UnmarshalCSV runs strings.ToUpper on the csv string
func (fName *Uppercase) UnmarshalCSV(csv string) error {
	fName.string = strings.ToUpper(csv)
	return nil
}

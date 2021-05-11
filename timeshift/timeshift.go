package timeshift

import (
	"time"

	"github.com/tkuchiki/go-timezone"
)

// DateTime structure.
type DateTime struct {
	time.Time
}

// MarshalCSV returns a time as as a string formatted to RFC3339.
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format(time.RFC3339), nil
}

// UnmarshalCSV accepts a string and unmarshals the string into a time.Time
// structure.
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	timeFormat := "1/2/06 3:04:05 PM"

	timeNameWest := "America/Los_Angeles"
	timeNameEast := "America/New_York"

	// Let's load the location we need for each timezone
	timeLocationWest, _ := time.LoadLocation(timeNameWest)
	timeLocationEast, _ := time.LoadLocation(timeNameEast)

	// Leveraging the go-timezone library
	tz := timezone.New()

	// Take the given CSV time, and parse it into a Pacific
	// Standard/Daylight Savings time zone.
	givenTime, err := time.ParseInLocation(timeFormat, csv, timeLocationWest)
	isDst := tz.IsDST(givenTime)

	// Since the go-timezone library looks things up via Abbrevations (e.g.
	// EST/PDT), we'll need to get the abbreviation for the two timezones
	// and make sure we're using Standard or Daylight Savings as well.
	timezoneAbbrWest, _ := tz.GetTimezoneAbbreviation(timeNameWest, isDst)
	timezoneAbbrEast, _ := tz.GetTimezoneAbbreviation(timeNameEast, isDst)

	// Once we have that, we can get the offset seconds from UTC and
	// leverage that number to add the number of seconds to jump
	// from Pacific to Eastern time.
	tzAbbrInfosWest, _ := tz.GetTzAbbreviationInfo(timezoneAbbrWest)
	tzAbbrInfosEast, _ := tz.GetTzAbbreviationInfo(timezoneAbbrEast)
	offsetInSeconds := tzAbbrInfosWest[0].Offset() - tzAbbrInfosEast[0].Offset()

	easternTime := givenTime.Add(time.Second * time.Duration(offsetInSeconds))

	date.Time = easternTime.In(timeLocationEast)
	return err
}

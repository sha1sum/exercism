package gigasecond

import "time"

const TestVersion = 2

// Add a gigasecond to a birthday and return a Time type
func AddGigasecond(birthday time.Time) time.Time {
	return birthday.Add(1e9 * time.Second)
}

// Get the time zone of the birthday
var location, err = time.LoadLocation("America/New_York")

// Set a birth date and time using the defined time zone
var Birthday time.Time = time.Date(1983, 9, 9, 9, 34, 0, 0, location)

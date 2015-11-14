package clock

import "fmt"

const TestVersion = 2

// Define the number of minutes in a day for use in calculations.
var minutesInDay int = 24 * 60

type Clock struct {
	// Hour of the day and minute of the hour.
	hour, minute int
}

// Return a Clock type from an hour of the day and minute of the hour
func Time(hour, minute int) Clock {
	return TimeByMinuteOfDay(hour * 60 + minute)
}

// Return a clock based on a number of minutes, either negative or positive.
func TimeByMinuteOfDay(minuteOfDay int) Clock {
	switch {
	// See if we've gone into following days...
	case minuteOfDay > minutesInDay:
		// ...and if so, get the remainder of minutes after dividing by days
		minuteOfDay %= minutesInDay
	// If we've gone into previous days...
	case minuteOfDay < 0:
		// ...we should get the remainder of minutes plus minutes in a day after dividing by negative days
		minuteOfDay = (minuteOfDay % (minutesInDay * -1)) + minutesInDay
	}
	if minuteOfDay == minutesInDay { minuteOfDay = 0 }
	// Return a new Clock type.
	return Clock{hour: minuteOfDay / 60, minute: minuteOfDay % 60}
}

// Format a Clock as a string
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add minutes to a Clock, and allow for negative minutes.
func (clock Clock) Add(minutes int) Clock {
	// Find the minute of the day (should be less than 1440 minutes)
	currentMinuteOfDay := clock.hour * 60 + clock.minute
	// Add the number of minutes to the minute of the day
	currentMinuteOfDay += minutes
	return TimeByMinuteOfDay(currentMinuteOfDay)
}
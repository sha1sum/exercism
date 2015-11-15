package clock

import "fmt"

const TestVersion = 2

// Define the number of minutes in a day for use in calculations.
const minutesInDay int = 24 * 60

// Define the number of minutes in a hour for use in calculations.
const minutesInHour int = 60

type Clock struct {
	// Minute of the day
	minute int
}

// Return a Clock type from an hour of the day and minute of the hour
func Time(hour, minute int) Clock {
	return TimeByMinuteOfDay(hour*minutesInHour + minute)
}

// Return a clock based on a number of minutes, either negative or positive.
func TimeByMinuteOfDay(minuteOfDay int) Clock {
	minuteOfDay %= minutesInDay

	if minuteOfDay < 0 {
		minuteOfDay += minutesInDay
	}

	return Clock{minuteOfDay}
}

// Format a Clock as a string
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minute/minutesInHour, clock.minute%minutesInHour)
}

// Add minutes to a Clock, and allow for negative minutes.
func (clock Clock) Add(minutes int) Clock {
	return TimeByMinuteOfDay(clock.minute + minutes)
}

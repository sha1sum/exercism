package leap

const TestVersion = 1

// Take a year argument and return whether that year is a leap year or not.
func IsLeapYear(year int) bool {
	// If the year is not divisible by 4, don't bother proceeding.
	if year % 4 != 0 { return false }
	switch {
	// If the year is divisible by 400, it's a leap year.
	case year % 400 == 0: return true
	// When the year is divisible by 100, but not by 400, it's not a leap year.
	case year % 100 == 0: return false
	// Every year that does not meet the first two cases is a leap year.
	default: return true
	}
}

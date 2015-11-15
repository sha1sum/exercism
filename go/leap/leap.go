package leap

const TestVersion = 1

// Take a year argument and return whether that year is a leap year or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && !(year%100 == 0 && !(year%400 == 0))
}

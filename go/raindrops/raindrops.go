package raindrops

import "fmt"

const TestVersion = 1

// Convert a number to raindrop speak
func Convert(number int) string {
	raindrop := ""
	if number%3 == 0 {
		raindrop += "Pling"
	}
	if number%5 == 0 {
		raindrop += "Plang"
	}
	if number%7 == 0 {
		raindrop += "Plong"
	}
	if len(raindrop) == 0 {
		raindrop = fmt.Sprintf("%d", number)
	}
	return raindrop
}

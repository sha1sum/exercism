package raindrops

import "fmt"

const TestVersion = 1

// Declare an empty struct to use for our map values
type Null struct {}

// Convert a number to raindrop speak
func Convert(number int) string {
	raindrop := ""
	// Get all of the factors for a number
	allFactors := factors(number)
	// Go through 3, 5, and 7, which we can translate, and add them to the translation if they're factors
	if contains(allFactors, 3) { raindrop += "Pling" }
	if contains(allFactors, 5) { raindrop += "Plang" }
	if contains(allFactors, 7) { raindrop += "Plong" }
	// If we haven't been able to translate anything so far, then we'll just have to pass the number
	if len(raindrop) == 0 { raindrop = fmt.Sprintf("%d", number) }
	return raindrop
}

// Get all of the factors for a given number and return them
func factors(number int) map[int]struct{} {
	// Use a map so we can easily tell if a value is present. The map's values are not important, just the keys.
	var allFactors = make(map[int]struct{})
	// Loop through all the integers from 1 through the number
	for i := 1; i <= number; i++ {
		// Add the factor to the map's keys if the number is evenly divisible by it.
		if number % i == 0 { allFactors[i] = Null{} }
	}
	return allFactors
}

// Check if our map contains a key for the integer in question
func contains(allFactors map[int]struct{}, number int) bool {
	_, ok := allFactors[number]
	// If the number is in the map keys, we're true
	if ok { return true }
	// Otherwise false
	return false
}
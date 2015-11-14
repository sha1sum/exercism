package diffsquares

// Get the square of the sums of 1 through the number provided
func SquareOfSums(number int) int {
	sum := number * (number + 1) / 2
	return sum * sum
}

// Get the sum of the squares from 1 through the number provided
func SumOfSquares(number int) int {
	return number * (number + 1) * (2 * number + 1) / 6
}

// Get the difference of the square of the sums and the sum of the squares of the number provided
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
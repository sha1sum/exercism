package diffsquares

// Get the square of the sums of 1 through the number provided
func SquareOfSums(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum * sum
}

// Get the sum of the squares from 1 through the number provided
func SumOfSquares(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i * i
	}
	return sum
}

// Get the difference of the square of the sums and the sum of the squares of the number provided
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
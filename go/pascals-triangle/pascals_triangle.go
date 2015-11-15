package pascal

// Build Pascal's triangle to number of specified rows.
func Triangle(rows int) [][]int {
	// Initialize the triangle with the first row being 1
	triangle := [][]int{[]int{1}}
	// If there's only one row, just return it
	if rows <= 1 {
		return triangle
	}
	// Loop through the number of rows, getting the previous row index
	for previousRow := 0; previousRow < rows-1; previousRow++ {
		// The first number in the row is always 1
		row := []int{1}
		// Loop through all the columns after the first and up to the last
		for i := 1; i < len(triangle[previousRow]); i++ {
			// Add the value of the column before the current one and the current column of the previous row and append
			row = append(row, triangle[previousRow][i-1]+triangle[previousRow][i])
		}
		// The last number in the row is always 1
		row = append(row, 1)
		// Append the row to the triangle
		triangle = append(triangle, row)
	}
	return triangle
}

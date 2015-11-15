package pascal

// Build Pascal's triangle. Formula: row! / (column! * (row - column)!)
func Triangle(rows int) [][]int {
	triangle := make([][]int, 0)
	// Return an empty slice if no rows
	if rows < 1 {
		return triangle
	}
	// The first row is always just 1
	triangle = append(triangle, []int{1})
	// For each row after the first
	for row := 1; row < rows; row++ {
		rowValues := make([]int, 0)
		// Calculate all of the descending natural numbers from the row number multiplied together
		rowDescendant := row
		for descendant := row - 1; descendant > 0; descendant-- {
			rowDescendant *= descendant
		}
		// For each column in the row
		for column := 0; column <= row; column++ {
			// First and last columns are always 1
			if column == 0 || column == row {
				rowValues = append(rowValues, 1)
				continue
			}
			// Calculate all of the descending natural numbers from the column number multiplied together
			columnDescendant := column
			for descendant := column - 1; descendant > 0; descendant-- {
				columnDescendant *= descendant
			}
			// Calculate all of the descending natural numbers from the difference between the
			// row and column numbers multiplied together
			rowMinusColumn := row - column
			rowColDescendant := rowMinusColumn
			for descendant := rowMinusColumn - 1; descendant > 0; descendant-- {
				rowColDescendant *= descendant
			}
			// Put all of the descendant multiplication values together into the final formula
			columnValue := rowDescendant / (columnDescendant * rowColDescendant)
			// Append the column value to the row
			rowValues = append(rowValues, columnValue)
		}
		// Append the row values to the triangle
		triangle = append(triangle, rowValues)
	}
	return triangle
}

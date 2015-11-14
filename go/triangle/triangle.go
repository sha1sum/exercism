package triangle

import "sort"

// Use integer values for kinds of triangles
type Kind int

// Definitions for the types of triangles
var Equ = Kind(0) // equilateral
var Iso = Kind(1) // isosceles
var Sca = Kind(2) // scalene
var NaT = Kind(3) // not a triangle

// Return a Kind of triangle from 3 lengths
func KindFromSides(a, b, c float64) Kind {
	// Put lengths into a slice and sort them in ascending order
	lengths := []float64{a, b, c}
	sort.Float64s(lengths)
	// If the shorter of the two sides are not greater than the largest side, or if each side's length is
	// not greater than 0, it's not a triangle.
	if lengths[0] + lengths[1] <= lengths[2] || !(a > 0 && b > 0 && c > 0) { return NaT }
	// If all sides are equal, it's equilateral
	if a == b && b == c { return Equ }
	// If two of the sides are the same, then it's isosceles
	if a == b || a == c || b == c { return Iso }
	// Otherwise it's got to be scalene
	return Sca
}
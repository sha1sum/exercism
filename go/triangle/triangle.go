package triangle

// Use integer values for kinds of triangles
type Kind int

// Types of triangles
const (
	Equ = iota
	Iso
	Sca
	NaT
)

// Return a Kind of triangle from 3 lengths
func KindFromSides(a, b, c float64) Kind {
	// If each two sides aren't greater than the other or any side is
	// not greater than 0, it's not a triangle.
	if !(a + b > c && a + c > b && b + c > a) { return NaT }
	// If all sides are equal, it's equilateral
	if a == b && b == c { return Equ }
	// If two of the sides are the same, then it's isosceles
	if a == b || a == c || b == c { return Iso }
	// Otherwise it's got to be scalene
	return Sca
}
package aocutils

// Vector2D represents a 2D set of ints.
type Vector2D struct {
	x, y int
}

// NewVector2D creates a new vector.
func NewVector2D(x, y int) Vector2D {
	return Vector2D{x, y}
}

// IsUnset returns whether the vector is set.
func (v Vector2D) IsUnset() bool {
	return v.x == 0 && v.y == 0
}

// X returns the x value of the vector.
func (v Vector2D) X() int {
	return v.x
}

// Y returns the y value of the vector.
func (v Vector2D) Y() int {
	return v.y
}

// Add adds another vector to the vector.
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

// Sub subtracts another vector to the vector.
func (v Vector2D) Sub(other Vector2D) Vector2D {
	return Vector2D{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

// Mul multiplies the vector with the values of another vector.
func (v Vector2D) Mul(factor int) Vector2D {
	return Vector2D{
		x: factor * v.x,
		y: factor * v.y,
	}
}

// Length calculates the length of the vector.
func (v Vector2D) Length() int {
	return v.x + v.y
}

// LengthSquared calculates the square of the length of the vector.
func (v Vector2D) LengthSquared() int {
	return v.x*v.x + v.y*v.y
}

// ManhattanLength calculates the Manhattan length of the vector.
func (v Vector2D) ManhattanLength() int {
	return AbsInt(v.x) + AbsInt(v.y)
}

// DistanceSquared calculates the difference between the squared length
// of this vector and another.
func (v Vector2D) DistanceSquared(o Vector2D) int {
	return v.Sub(o).LengthSquared()
}

// ManhattanDistance calculates the Manhattan distance between the vector
// and another vector.
func (v Vector2D) ManhattanDistance(o Vector2D) int {
	return v.Sub(o).ManhattanLength()
}

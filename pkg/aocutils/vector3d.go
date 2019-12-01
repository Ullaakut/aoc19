package aocutils

// Vector2D represents a 2D set of ints.
type Vector3D struct {
	x, y, z int
}

// Add adds another vector to the vector.
func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

// Sub subtracts another vector to the vector.
func (v Vector3D) Sub(other Vector3D) Vector3D {
	return Vector3D{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

// Mul multiplies the vector with the values of another vector.
func (v Vector3D) Mul(factor int) Vector3D {
	return Vector3D{
		x: factor * v.x,
		y: factor * v.y,
		z: factor * v.z,
	}
}

// Length calculates the length of the vector.
func (v Vector3D) Length() int {
	return v.x + v.y + v.z
}

// LengthSquared calculates the square of the length of the vector.
func (v Vector3D) LengthSquared() int {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// ManhattanLength calculates the Manhattan length of the vector.
func (v Vector3D) ManhattanLength() int {
	return AbsInt(v.x) + AbsInt(v.y) + AbsInt(v.z)
}

// DistanceSquared calculates the difference between the squared length
// of this vector and another.
func (v Vector3D) DistanceSquared(o Vector3D) int {
	return v.Sub(o).LengthSquared()
}

// ManhattanDistance calculates the Manhattan distance between the vector
// and another vector.
func (v Vector3D) ManhattanDistance(o Vector3D) int {
	return v.Sub(o).ManhattanLength()
}

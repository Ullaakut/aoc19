package aocutils

type Vector3D struct {
	x, y, z int
}

func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector3D) Sub(other Vector3D) Vector3D {
	return Vector3D{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3D) Mul(factor int) Vector3D {
	return Vector3D{
		x: factor * v.x,
		y: factor * v.y,
		z: factor * v.z,
	}
}

func (v Vector3D) LengthSquared() int {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector3D) ManhattanLength() int {
	return AbsInt(v.x) + AbsInt(v.y) + AbsInt(v.z)
}

func (v Vector3D) DistanceSquared(o Vector3D) int {
	return v.Sub(o).LengthSquared()
}

func (v Vector3D) ManhattanDistance(o Vector3D) int {
	return v.Sub(o).ManhattanLength()
}

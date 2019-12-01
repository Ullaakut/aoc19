package aocutils

type Vector2D struct {
	x, y int
}

func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector2D) Sub(other Vector2D) Vector2D {
	return Vector2D{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector2D) Mul(factor int) Vector2D {
	return Vector2D{
		x: factor * v.x,
		y: factor * v.y,
	}
}

func (v Vector2D) LengthSquared() int {
	return v.x*v.x + v.y*v.y
}

func (v Vector2D) ManhattanLength() int {
	return AbsInt(v.x) + AbsInt(v.y)
}

func (v Vector2D) DistanceSquared(o Vector2D) int {
	return v.Sub(o).LengthSquared()
}

func (v Vector2D) ManhattanDistance(o Vector2D) int {
	return v.Sub(o).ManhattanLength()
}

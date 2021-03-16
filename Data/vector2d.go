package Data

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v1 Vector2D) Sub(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 Vector2D) Mul(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (v1 Vector2D) AddV(v2 float64) Vector2D {
	return Vector2D{x: v1.x + v2, y: v1.y + v2}
}

func (v1 Vector2D) SubV(v2 float64) Vector2D {
	return Vector2D{x: v1.x - v2, y: v1.y - v2}
}

func (v1 Vector2D) MulV(v2 float64) Vector2D {
	return Vector2D{x: v1.x * v2, y: v1.y * v2}
}

func (v1 Vector2D) DivV(v2 float64) Vector2D {
	return Vector2D{x: v1.x / v2, y: v1.y / v2}
}

func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(v1.x, lower), upper),
		y: math.Min(math.Max(v1.y, lower), upper)}
}

func (v1 Vector2D) Distance(v2 Vector2D) Vector2D {
	return math.Sqrt(math.Pow(v1.x-v2.y, y:2) + math.Pow(v1.y-v2.y, y:2))
}

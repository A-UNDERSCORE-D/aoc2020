package vector

var (
	V2Up          = Vec2d{X: 0, Y: 1}
	V2Down        = Vec2d{X: 0, Y: -1}
	V2Left        = Vec2d{X: -1, Y: 0}
	V2Right       = Vec2d{X: 1, Y: 0}
	V2BottomLeft  = Vec2d{X: -1, Y: -1}
	V2BottomRight = Vec2d{X: 1, Y: -1}
	V2UpRight     = Vec2d{X: 1, Y: 1}
	V2UpLeft      = Vec2d{X: -1, Y: 1}

	V2Directions = [...]Vec2d{V2Up, V2Down, V2Left, V2Right, V2BottomLeft, V2BottomRight, V2UpLeft, V2UpRight}
)

type Vec2d struct {
	X int
	Y int
}

func New2d(x, y int) Vec2d {
	return Vec2d{X: x, Y: y}
}

func (p Vec2d) Add(other Vec2d) Vec2d {
	return Vec2d{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Vec2d) AddInt(i int) Vec2d {
	return Vec2d{
		X: p.X + i,
		Y: p.Y + i,
	}
}

func (p Vec2d) Sub(other Vec2d) Vec2d {
	return Vec2d{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Vec2d) SubInt(i int) Vec2d {
	return Vec2d{
		X: p.X - i,
		Y: p.Y - i,
	}
}

func (p Vec2d) Mul(other Vec2d) Vec2d {
	return Vec2d{
		X: p.X * other.X,
		Y: p.Y * other.Y,
	}
}

func (p Vec2d) MulInt(i int) Vec2d {
	return Vec2d{X: p.X * i, Y: p.Y * i}
}

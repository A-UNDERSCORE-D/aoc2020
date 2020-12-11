package vector

var (
	V2Up          = Vec2d{x: 0, y: 1}
	V2Down        = Vec2d{x: 0, y: -1}
	V2Left        = Vec2d{x: -1, y: 0}
	V2Right       = Vec2d{x: 1, y: 0}
	V2BottomLeft  = Vec2d{x: -1, y: -1}
	V2BottomRight = Vec2d{x: 1, y: -1}
	V2UpRight     = Vec2d{x: 1, y: 1}
	V2UpLeft      = Vec2d{x: -1, y: 1}

	V2Directions = [...]Vec2d{V2Up, V2Down, V2Left, V2Right, V2BottomLeft, V2BottomRight, V2UpLeft, V2UpRight}
)

type Vec2d struct {
	x int
	y int
}

func New2d(x, y int) Vec2d {
	return Vec2d{x: x, y: y}
}

func (p Vec2d) Add(other Vec2d) Vec2d {
	return Vec2d{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p Vec2d) Sub(other Vec2d) Vec2d {
	return Vec2d{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}

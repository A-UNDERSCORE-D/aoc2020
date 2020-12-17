package vector

var (
	V3Up          = Vec3d{X: 0, Y: 1, Z: 0}
	V3Down        = Vec3d{X: 0, Y: -1, Z: 0}
	V3Left        = Vec3d{X: -1, Y: 0, Z: 0}
	V3Right       = Vec3d{X: 1, Y: 0, Z: 0}
	V3BottomLeft  = Vec3d{X: -1, Y: -1, Z: 0}
	V3BottomRight = Vec3d{X: 1, Y: -1, Z: 0}
	V3UpperRight  = Vec3d{X: 1, Y: 1, Z: 0}
	V3UpperLeft   = Vec3d{X: -1, Y: 1, Z: 0}

	V3UpUp          = Vec3d{X: 0, Y: 1, Z: 1}
	V3UpDown        = Vec3d{X: 0, Y: -1, Z: 1}
	V3UpLeft        = Vec3d{X: -1, Y: 0, Z: 1}
	V3UpRight       = Vec3d{X: 1, Y: 0, Z: 1}
	V3UpBottomLeft  = Vec3d{X: -1, Y: -1, Z: 1}
	V3UpBottomRight = Vec3d{X: 1, Y: -1, Z: 1}
	V3UpUpperRight  = Vec3d{X: 1, Y: 1, Z: 1}
	V3UpUpperLeft   = Vec3d{X: -1, Y: 1, Z: 1}

	V3DownUp          = Vec3d{X: 0, Y: 1, Z: -1}
	V3DownDown        = Vec3d{X: 0, Y: -1, Z: -1}
	V3DownLeft        = Vec3d{X: -1, Y: 0, Z: -1}
	V3DownRight       = Vec3d{X: 1, Y: 0, Z: -1}
	V3DownBottomLeft  = Vec3d{X: -1, Y: -1, Z: -1}
	V3DownBottomRight = Vec3d{X: 1, Y: -1, Z: -1}
	V3DownUpperRight  = Vec3d{X: 1, Y: 1, Z: -1}
	V3DownUpperLeft   = Vec3d{X: -1, Y: 1, Z: -1}

	V3Neighbours = []Vec3d{
		{X: 0, Y: 0, Z: 1},
		{X: 0, Y: 1, Z: 1},
		{X: 0, Y: -1, Z: 1},
		{X: -1, Y: 0, Z: 1},
		{X: 1, Y: 0, Z: 1},
		{X: -1, Y: -1, Z: 1},
		{X: 1, Y: -1, Z: 1},
		{X: 1, Y: 1, Z: 1},
		{X: -1, Y: 1, Z: 1},

		{X: 0, Y: 1, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: -1, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 0},
		{X: -1, Y: -1, Z: 0},
		{X: 1, Y: -1, Z: 0},
		{X: 1, Y: 1, Z: 0},
		{X: -1, Y: 1, Z: 0},

		{X: 0, Y: 0, Z: -1},
		{X: 0, Y: 1, Z: -1},
		{X: 0, Y: -1, Z: -1},
		{X: -1, Y: 0, Z: -1},
		{X: 1, Y: 0, Z: -1},
		{X: -1, Y: -1, Z: -1},
		{X: 1, Y: -1, Z: -1},
		{X: 1, Y: 1, Z: -1},
		{X: -1, Y: 1, Z: -1},
	}
)

type Vec3d struct {
	X int
	Y int
	Z int
}

func New3d(x, y, z int) Vec3d {
	return Vec3d{X: x, Y: y, Z: z}
}

func (p Vec3d) Add(other Vec3d) Vec3d {
	return Vec3d{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
	}
}

func (p Vec3d) AddInt(i int) Vec3d {
	return Vec3d{
		X: p.X + i,
		Y: p.Y + i,
		Z: p.Z + i,
	}
}

func (p Vec3d) Sub(other Vec3d) Vec3d {
	return Vec3d{
		X: p.X - other.X,
		Y: p.Y - other.Y,
		Z: p.Z - other.Z,
	}
}

func (p Vec3d) SubInt(i int) Vec3d {
	return Vec3d{
		X: p.X - i,
		Y: p.Y - i,
		Z: p.Z - i,
	}
}

func (p Vec3d) Mul(other Vec3d) Vec3d {
	return Vec3d{
		X: p.X * other.X,
		Y: p.Y * other.Y,
		Z: p.Z * other.Z,
	}
}

func (p Vec3d) MulInt(i int) Vec3d {
	return Vec3d{X: p.X * i, Y: p.Y * i, Z: p.Z * i}
}

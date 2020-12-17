package vector

var V4Neighbours []Vec4d = func() (out []Vec4d) {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				for w := -1; w < 2; w++ {
					target := New4d(x, y, z, w)
					if target == (Vec4d{0, 0, 0, 0}) {
						continue
					}
					out = append(out, target)
				}
			}
		}
	}

	return
}()

type Vec4d struct {
	X, Y, Z, W int
}

func New4d(x, y, z, w int) Vec4d {
	return Vec4d{X: x, Y: y, Z: z, W: w}
}

func (p Vec4d) Add(other Vec4d) Vec4d {
	return Vec4d{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
		W: p.W + other.W,
	}
}

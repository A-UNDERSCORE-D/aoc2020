package draw

import (
	"image"
	"image/color"
)

type SquareDot struct {
	Colour   color.Color
	Position image.Point
}

func (d *SquareDot) Render(i *image.Paletted, offset, size int) {
	startPos := d.Position.Mul(offset)
	for y := startPos.Y; y < startPos.Y+size; y++ {
		for x := startPos.X; x < startPos.X+size; x++ {
			i.Set(x, y, d.Colour)
		}
	}
}

func (d *SquareDot) GetPosition() image.Point {
	return d.Position
}

type XDot struct {
	SquareDot
	BackgroundColour color.Color
	Width            int
}

func (d *XDot) isWithin(a, b int) bool {
	return a-b <= d.Width && a-b >= 0
}

func (d *XDot) Render(i *image.Paletted, offset, size int) {
	startPos := d.Position.Mul(offset)
	xLeft := 0
	xRight := size - 1

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			colour := d.BackgroundColour
			if d.isWithin(x, xLeft) || d.isWithin(x, xRight) {
				colour = color.White
			}
			i.Set(x+startPos.X, y+startPos.Y, colour)
		}
		xLeft++
		xRight--
	}
}

type Dot interface {
	Render(i *image.Paletted, offset, size int)
}

func RenderDots(maxX, maxY int, palette color.Palette, dots []Dot, dotSize, dotSpace int) *image.Paletted {
	offset := dotSize + dotSpace
	img := image.NewPaletted(image.Rect(-dotSpace, -dotSpace, maxX*offset, maxY*offset), palette)
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			img.Set(x, y, color.Black)
		}
	}

	for _, d := range dots {
		d.Render(img, offset, dotSize)
	}
	return img
}

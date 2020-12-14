package draw

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"os"
	"strings"
	"testing"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const td = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func Test(t *testing.T) {
	var dots []Dot
	lines := util.ReadLines("../../2020/11/input.txt")
	lines = strings.Split(td, "\n")
	maxX := 0
	maxY := 0
	for y, row := range lines {
		maxY = y
		for x, r := range row {
			maxX = x
			var dot Dot
			switch r {
			case '.':
				dot = &SquareDot{Colour: color.White, Position: image.Point{x, y}}
			case 'L':
				dot = &XDot{
					SquareDot:        SquareDot{Colour: color.RGBA{R: 0xFF, A: 0xFF}, Position: image.Point{x, y}},
					BackgroundColour: color.RGBA{B: 0xFF, A: 0xFF},
					Width:            3,
				}
			case '#':
				dot = &SquareDot{Colour: color.RGBA{R: 0xFF, A: 0xFF}, Position: image.Point{x, y}}

			}
			dots = append(dots, dot)
		}
	}

	img := RenderDots(maxX, maxY, palette.Plan9, dots, 50, 10)
	f, err := os.Create("./out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t.Log(png.Encode(f, image.Image(img)))
}

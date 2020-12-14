package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"awesome-dragon.science/go/adventofcode2020/util/vector"
)

const testData = `F10
N3
F7
R90
F11`

func main() {
	input := util.ReadLines("input.txt")
	_ = strings.Split(testData, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

var dirLut = map[rune]vector.Vec2d{
	'N': vector.V2Up,
	'S': vector.V2Down,
	'E': vector.V2Right,
	'W': vector.V2Left,
}

var dirLut2 = []rune{'N', 'E', 'S', 'W'}

func part1(input []string) string {
	curPos := vector.Vec2d{}
	currentFacing := 'E'
	currentDirIdx := 1

	moves := [][2]image.Point{}

	for _, ins := range input {
		insChr := rune(ins[0])
		num := util.GetInt(ins[1:])
		switch insChr {
		case 'F':
			oldPos := curPos
			curPos = curPos.Add(dirLut[dirLut2[currentDirIdx]].MulInt(num))
			moves = append(moves, [2]image.Point{image.Point(oldPos), image.Point(curPos)})
		case 'N', 'E', 'S', 'W':
			oldPos := curPos
			curPos = curPos.Add(dirLut[insChr].MulInt(num))
			moves = append(moves, [2]image.Point{image.Point(oldPos), image.Point(curPos)})
		case 'L':
			// Because I cant math when Im tired, apparently
			for i := 0; i < num/90; i++ {
				currentDirIdx--
				if currentDirIdx == -1 {
					currentDirIdx = 3
				}
			}
			currentFacing = dirLut2[currentDirIdx]
		case 'R':
			currentDirIdx = ((num / 90) + currentDirIdx) % 4
			currentFacing = dirLut2[currentDirIdx]
		}
	}

	minX := math.MaxInt64
	maxX := math.MinInt64
	minY := math.MaxInt64
	maxY := math.MinInt64
	for _, v := range moves {
		for _, p := range v {
			minX = util.Min(minX, p.X)
			maxX = util.Max(maxX, p.X)
			minY = util.Min(minY, p.Y)
			maxY = util.Max(maxY, p.Y)
		}
	}

	img := image.NewRGBA(image.Rect(minX-1, minY-1, maxX+1, maxY+1))
	// one := image.Pt(1, 1)
	for i, v := range moves {
		const offset = 1
		p1 := vector.Vec2d(v[0]).AddInt(offset)
		p2 := vector.Vec2d(v[1]).AddInt(offset)
		// if util.Max(p1.X, p2.X) == p2.X {
		// 	p1 = v[1]
		// 	p2 = v[0]
		// }

		rect := image.Rect(
			util.Min(p1.X, p2.X),
			util.Min(p1.Y, p2.Y),
			util.Max(p1.X, p2.X)+1,
			util.Max(p1.Y, p2.Y)+1,
		)
		draw.Draw(img, rect, &image.Uniform{color.NRGBA64{R:  math.MaxUint16, A: math.MaxUint16, G: uint16(i)*1000, B:  uint16(i)*100}}, image.Point{}, draw.Src)
	}
	// draw.Draw(img, image.Rectangle{image.Point(oldPos.SubInt(1)), image.Point(curPos.AddInt(1))}, &image.Uniform{color.White}, image.Point{}, draw.Src)
	_ = currentFacing
	util.PNGQuick("./test.png", img)
	return strconv.FormatInt(int64(util.Abs(curPos.X)+util.Abs(curPos.Y)), 10)
}

func part2(input []string) string {
	shipPos := vector.Vec2d{}
	waypointPos := vector.Vec2d{X: 10, Y: 1}

	for _, ins := range input {
		insChr := rune(ins[0])
		num := util.GetInt(ins[1:])
		switch insChr {
		case 'F':
			shipPos = shipPos.Add(waypointPos.MulInt(num))
		case 'N', 'E', 'S', 'W':
			waypointPos = waypointPos.Add(dirLut[insChr].MulInt(num))
		case 'L':
			for i := 0; i < num/90; i++ {
				waypointPos = vector.New2d(-waypointPos.Y, waypointPos.X)
			}
		case 'R':
			for i := 0; i < num/90; i++ {
				waypointPos = vector.New2d(waypointPos.Y, -waypointPos.X)
			}
		}
	}

	return strconv.FormatInt(int64(util.Abs(shipPos.X)+util.Abs(shipPos.Y)), 10)
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"awesome-dragon.science/go/adventofcode2020/util/draw"
	"awesome-dragon.science/go/adventofcode2020/util/vector"
)

const testData = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func main() {
	input := util.ReadLines("input.txt")
	// input = strings.Split(testData, "\n")
	// f, _ := os.Create("profile")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	startTime := time.Now()
	parsed := parseField(input)
	fmt.Println("Parsing done. Took: ", time.Since(startTime))
	startTime = time.Now()
	res := part1(parsed)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part2(parsed)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

const (
	unknown = iota
	emptySeat
	fullSeat
	floor
)

const (
	dotSize  = 10
	dotSpace = 2
)

type position struct {
	vector.Vec2d
	state int
}

type field map[vector.Vec2d]position

func parseField(fieldStr []string) field {
	out := make(map[vector.Vec2d]position, len(fieldStr)*len(fieldStr[0]))

	for y, row := range fieldStr {
		for x, seat := range row {
			toSet := position{Vec2d: vector.New2d(x, y)}
			switch seat {
			case '.':
				toSet.state = floor
			case 'L':
				toSet.state = emptySeat
			case '#':
				toSet.state = fullSeat
			default:
				fmt.Println("UNKNOWN SEAT TYPE: ", x, y)
			}
			out[toSet.Vec2d] = toSet
		}
	}
	return out
}

func (f field) toImage(dotSize, spacing int) image.Image {
	dots := []draw.Dot{}
	f.iter(func(p position) {
		var dot draw.Dot
		switch p.state {
		case floor:
			dot = &draw.SquareDot{Colour: color.White, Position: image.Point(p.Vec2d)}
		case emptySeat:
			dot = &draw.SquareDot{Colour: color.RGBA{B: 0xFF, A: 0xFF}, Position: image.Point(p.Vec2d)}
		case fullSeat:
			dot = &draw.SquareDot{Colour: color.RGBA{R: 0xFF, A: 0xFF}, Position: image.Point(p.Vec2d)}
		}
		dots = append(dots, dot)
	})

	return draw.RenderDots(f.maxX(), f.maxY(), palette.Plan9, dots, dotSize, spacing)
}

var maxX = -1

func (f field) maxX() int {
	if maxX != -1 {
		return maxX
	}
	for i := 0; ; i++ {
		if _, ok := f[vector.New2d(i, 0)]; !ok {
			maxX = i
			return i
		}
	}
}

var maxY = -1

func (f field) maxY() int {
	if maxY != -1 {
		return maxY
	}
	for i := 0; ; i++ {
		if _, ok := f[vector.New2d(0, i)]; !ok {
			maxY = i
			return i
		}
	}
}

func (f field) String() string {
	lineBuilder := strings.Builder{}
	for y := 0; y < f.maxY(); y++ {
		for x := 0; x < f.maxX(); x++ {
			switch f[vector.New2d(x, y)].state {
			case floor:
				lineBuilder.WriteRune('.')
			case fullSeat:
				lineBuilder.WriteRune('#')
			case emptySeat:
				lineBuilder.WriteRune('L')
			case unknown:
				lineBuilder.WriteRune('?')
			}
		}
		lineBuilder.WriteRune('\n')
	}

	return lineBuilder.String()
}

func (f field) nextState(pos vector.Vec2d, part1 bool) position {
	fullAround := 0
	currentSeat, ok := f[pos]
	if !ok {
		panic("Bad seat passed")
	}

	if currentSeat.state == floor {
		return currentSeat
	}

	for _, direction := range vector.V2Directions {
		var other position
		var ok bool
		if part1 {
			other, ok = f[pos.Add(direction)]
		} else {
			var otherPos vector.Vec2d
			otherPos, ok = f.rayCastFindNextSeat(pos, direction)
			other = f[otherPos]
		}

		if !ok {
			continue // Must be at the edges
		}

		if other.state == fullSeat {
			fullAround++
		}
	}

	switch {
	case currentSeat.state == emptySeat && fullAround == 0:
		currentSeat.state = fullSeat
		return currentSeat
	case currentSeat.state == fullSeat && (part1 && fullAround >= 4) || (!part1 && fullAround >= 5):
		currentSeat.state = emptySeat
		return currentSeat
	}

	return currentSeat
}

func (f field) iter(fn func(position)) {
	for y := 0; y < f.maxY(); y++ {
		for x := 0; x < f.maxX(); x++ {
			fn(f[vector.New2d(x, y)])
		}
	}
}

func (f field) stepPart1() field {
	out := make(field, len(f))

	f.iter(func(p position) { out[p.Vec2d] = f.nextState(p.Vec2d, true) })

	return out
}

func (f field) Equals(other field) bool {
	if len(f) != len(other) {
		return false
	}

	for k, v := range f {
		if other[k] != v {
			return false
		}
	}

	return true
}

func part1(input field) string {
	// var images []image.Image
	oldField := input
	// images = append(images, oldField.toImage(dotSize, dotSpace))
	var final field
	for {
		newField := oldField.stepPart1()
		if oldField.Equals(newField) {
			final = newField
			break
		}
		// images = append(images, newField.toImage(dotSize, dotSpace))
		// fmt.Print("\033[2J", newField.String())
		oldField = newField

		// time.Sleep(time.Millisecond * 150)
	}
	seatCount := 0
	final.iter(func(p position) {
		if p.state == fullSeat {
			seatCount++
		}
	})

	// finalImg := final.toImage(dotSize, dotSpace)
	// images = append(images, finalImg)

	// os.Mkdir("./p1", 0o744)
	// for i, v := range images {
	// 	f, _ := os.Create(fmt.Sprintf("./p1/%d.png", i))
	// 	if err := png.Encode(f, v); err != nil {
	// 		panic(err)
	// 	}
	// 	f.Close()
	// }
	return strconv.FormatInt(int64(seatCount), 10)
}

func (f field) rayCastFindNextSeat(pos vector.Vec2d, direction vector.Vec2d) (vector.Vec2d, bool) {
	lastPos := pos
	for {
		nextPos := lastPos.Add(direction)
		seat, ok := f[nextPos]
		if !ok {
			return pos, false
		}

		if seat.state == fullSeat || seat.state == emptySeat {
			return seat.Vec2d, true
		}
		lastPos = nextPos
	}
}

func (f field) stepPart2() field {
	out := make(field, len(f))
	f.iter(func(p position) { out[p.Vec2d] = f.nextState(p.Vec2d, false) })
	return out
}

func part2(input field) string {
	// var images []*image.Paletted
	oldField := input
	// images = append(images, oldField.toImage(dotSize, dotSpace).(*image.Paletted))
	var final field
	for {
		newField := oldField.stepPart2()
		if oldField.Equals(newField) {
			final = newField
			break
		}
		// images = append(images, newField.toImage(dotSize, dotSpace).(*image.Paletted))
		// fmt.Print("\033[2J", newField.String())
		oldField = newField

		// time.Sleep(time.Millisecond * 150)
	}
	seatCount := 0
	final.iter(func(p position) {
		if p.state == fullSeat {
			seatCount++
		}
	})
	// finalImg := final.toImage(dotSize, dotSpace).(*image.Paletted)
	// images = append(images, finalImg)
	// os.Mkdir("./p2", 0o744)
	// for i, v := range images {
	// 	f, _ := os.Create(fmt.Sprintf("./p2/%d.png", i))
	// 	if err := png.Encode(f, v); err != nil {
	// 		panic(err)
	// 	}
	// 	f.Close()
	// }
	return strconv.FormatInt(int64(seatCount), 10)
}

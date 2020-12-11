package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
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

	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))

	time.Sleep(time.Second * 5)
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

const (
	unknown = iota
	emptySeat
	fullSeat
	floor
)

type position struct {
	vector.Vec2d
	state int
}

type field map[vector.Vec2d]position

func parseField(fieldStr []string) field {
	out := make(map[vector.Vec2d]position)

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

func (f field) nextStatePart1(pos vector.Vec2d) position {
	emptyAround := 0
	fullAround := 0
	floorAround := 0
	currentSeat, ok := f[pos]
	if !ok {
		panic("Bad seat passed")
	}

	if currentSeat.state == floor {
		return currentSeat
	}

	for _, direction := range vector.V2Directions {
		other, ok := f[pos.Add(direction)]
		if !ok {
			continue // Must be at the edges
		}

		switch other.state {
		case unknown:
			fmt.Println("UNKNOWN STATE AT ", pos.Add(direction))
		case fullSeat:
			fullAround++
		case emptySeat:
			emptyAround++
		case floor:
			floorAround++
		}
	}

	switch {
	case currentSeat.state == emptySeat && fullAround == 0:
		currentSeat.state = fullSeat
		return currentSeat
	case currentSeat.state == fullSeat && fullAround >= 4:
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
	out := make(field)
	for y := 0; y < f.maxY(); y++ {
		for x := 0; x < f.maxX(); x++ {
			targetPos := vector.New2d(x, y)
			out[targetPos] = f.nextStatePart1(targetPos)
		}
	}

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

func part1(input []string) string {
	parsed := parseField(input)
	oldField := parsed
	var final field
	for {
		newField := oldField.stepPart1()
		if oldField.Equals(newField) {
			final = newField
			break
		}
		fmt.Print("\033[2J", newField.String())
		oldField = newField

		time.Sleep(time.Millisecond * 150)
	}
	seatCount := 0
	final.iter(func(p position) {
		if p.state == fullSeat {
			seatCount++
		}
	})

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

func (f field) nextStatePart2(pos vector.Vec2d) position {
	emptyAround := 0
	fullAround := 0
	floorAround := 0
	currentSeat, ok := f[pos]
	if !ok {
		panic("Bad seat passed")
	}

	if currentSeat.state == floor {
		return currentSeat
	}

	for _, direction := range vector.V2Directions {
		otherPos, ok := f.rayCastFindNextSeat(pos, direction)
		if !ok {
			continue // Must be at the edges
		}

		other := f[otherPos]

		switch other.state {
		case unknown:
			fmt.Println("UNKNOWN STATE AT ", pos.Add(direction))
		case fullSeat:
			fullAround++
		case emptySeat:
			emptyAround++
		case floor:
			floorAround++
		}
	}

	switch {
	case currentSeat.state == emptySeat && fullAround == 0:
		currentSeat.state = fullSeat
		return currentSeat
	case currentSeat.state == fullSeat && fullAround >= 5:
		currentSeat.state = emptySeat
		return currentSeat
	}

	return currentSeat
}

func (f field) stepPart2() field {
	out := make(field)

	f.iter(func(p position) {
		out[p.Vec2d] = f.nextStatePart2(p.Vec2d)
	})
	return out
}

func part2(input []string) string {
	parsed := parseField(input)
	oldField := parsed
	var final field
	for {
		newField := oldField.stepPart2()
		if oldField.Equals(newField) {
			final = newField
			break
		}
		fmt.Print("\033[2J", newField.String())
		oldField = newField

		time.Sleep(time.Millisecond * 150)
	}
	seatCount := 0
	final.iter(func(p position) {
		if p.state == fullSeat {
			seatCount++
		}
	})

	return strconv.FormatInt(int64(seatCount), 10)
}

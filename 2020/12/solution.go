package main

import (
	"fmt"
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

	for _, ins := range input {
		insChr := rune(ins[0])
		num := util.GetInt(ins[1:])
		switch insChr {
		case 'F':
			curPos = curPos.Add(dirLut[dirLut2[currentDirIdx]].MulInt(num))
		case 'N', 'E', 'S', 'W':
			curPos = curPos.Add(dirLut[insChr].MulInt(num))
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

	_ = currentFacing
	fmt.Println(curPos)
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

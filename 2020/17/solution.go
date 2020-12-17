package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"awesome-dragon.science/go/adventofcode2020/util/vector"
)

const testData = `
.#.
..#
###`

func main() {
	input := util.ReadLines("input.txt")
	_ = strings.Split(testData[1:], "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func cloneMap(in map[vector.Vec3d]rune) map[vector.Vec3d]rune {
	out := make(map[vector.Vec3d]rune, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

const (
	unknown  rune = 0x0
	active   rune = '#'
	inactive rune = '.'
)

func fieldExtents(field map[vector.Vec3d]rune) (maxX, maxY, maxZ, minX, minY, minZ int) {
	maxX = math.MinInt64
	maxY = math.MinInt64
	maxZ = math.MinInt64
	minX = math.MaxInt64
	minY = math.MaxInt64
	minZ = math.MaxInt64

	for k := range field {
		maxX = util.Max(maxX, k.X)
		maxY = util.Max(maxY, k.Y)
		maxZ = util.Max(maxZ, k.Z)

		minX = util.Min(minX, k.X)
		minY = util.Min(minY, k.Y)
		minZ = util.Min(minZ, k.Z)
	}

	return
}

func printField(field map[vector.Vec3d]rune) {
	out := strings.Builder{}
	maxX, maxY, maxZ, minX, minY, minZ := fieldExtents(field)
	for z := minZ; z <= maxZ; z++ {
		out.WriteString(fmt.Sprintf("Z: %d\n", z))
		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				state := field[vector.New3d(x, y, z)]
				switch state {
				case active, inactive:
					out.WriteRune(state)
				case unknown:
					out.WriteRune(inactive)
				}
			}
			out.WriteRune('\n')
		}
		out.WriteRune('\n')
	}
	fmt.Println(out.String())
}

func getAllPossiblePositions(field map[vector.Vec3d]rune) <-chan vector.Vec3d {
	out := make(chan vector.Vec3d)
	go func() {
		maxX, maxY, maxZ, minX, minY, minZ := fieldExtents(field)
		for z := minZ - 1; z <= maxZ+1; z++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for x := minX - 1; x <= maxX+1; x++ {
					out <- vector.New3d(x, y, z)
				}
			}
		}

		close(out)
	}()
	return out
}

func step(field map[vector.Vec3d]rune, print bool) map[vector.Vec3d]rune {
	if print {
		printField(field)
	}
	newField := cloneMap(field)
	for pos := range getAllPossiblePositions(field) {
		posState := field[pos]
		activeCount := 0
		for _, dir := range vector.V3Neighbours {
			targetPos := pos.Add(dir)
			neighbour := field[targetPos]
			if neighbour == active {
				activeCount++
			}
		}

		switch posState {
		case active:
			if activeCount != 2 && activeCount != 3 {
				newField[pos] = inactive
			}
		case inactive, unknown:
			if activeCount == 3 {
				newField[pos] = active
			}
		}
	}
	return newField
}

func part1(input []string) string {
	field := make(map[vector.Vec3d]rune)
	for y, line := range input {
		for x, r := range line {
			field[vector.New3d(x, y, 0)] = r
			field[vector.New3d(x, y, 1)] = inactive
			field[vector.New3d(x, y, -1)] = inactive
		}
	}

	for i := 0; i < 6; i++ {
		field = step(field, false)
	}
	count := 0
	for _, v := range field {
		if v == active {
			count++
		}
	}

	return fmt.Sprint(count)
}

func cloneMap4(in map[vector.Vec4d]rune) map[vector.Vec4d]rune {
	out := make(map[vector.Vec4d]rune, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func fieldExtents4(field map[vector.Vec4d]rune) (maxX, maxY, maxZ, maxW, minX, minY, minZ, minW int) {
	maxX = math.MinInt64
	maxY = math.MinInt64
	maxZ = math.MinInt64
	maxW = math.MinInt64

	minX = math.MaxInt64
	minY = math.MaxInt64
	minZ = math.MaxInt64
	minW = math.MaxInt64

	for k := range field {
		maxX = util.Max(maxX, k.X)
		maxY = util.Max(maxY, k.Y)
		maxZ = util.Max(maxZ, k.Z)
		maxW = util.Max(maxW, k.W)

		minX = util.Min(minX, k.X)
		minY = util.Min(minY, k.Y)
		minZ = util.Min(minZ, k.Z)
		minW = util.Min(minZ, k.W)
	}

	return
}

func getAllPossiblePositions4(field map[vector.Vec4d]rune) <-chan vector.Vec4d {
	out := make(chan vector.Vec4d)
	go func() {
		maxX, maxY, maxZ, maxW, minX, minY, minZ, minW := fieldExtents4(field)
		for z := minZ - 1; z <= maxZ+1; z++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for x := minX - 1; x <= maxX+1; x++ {
					for w := minW - 1; w <= maxW+1; w++ {
						out <- vector.New4d(x, y, z, w)
					}
				}
			}
		}

		close(out)
	}()
	return out
}

func step4(field map[vector.Vec4d]rune, print bool) map[vector.Vec4d]rune {
	newField := cloneMap4(field)
	for pos := range getAllPossiblePositions4(field) {
		posState := field[pos]
		activeCount := 0
		for _, dir := range vector.V4Neighbours {
			targetPos := pos.Add(dir)
			neighbour := field[targetPos]
			if neighbour == active {
				activeCount++
			}
		}

		switch posState {
		case active:
			if activeCount != 2 && activeCount != 3 {
				newField[pos] = inactive
			}
		case inactive, unknown:
			if activeCount == 3 {
				newField[pos] = active
			}
		}
	}
	return newField
}

func part2(input []string) string {
	field := make(map[vector.Vec4d]rune)
	for y, line := range input {
		for x, r := range line {
			field[vector.New4d(x, y, 0, 0)] = r

			field[vector.New4d(x, y, 1, 1)] = inactive
			field[vector.New4d(x, y, 1, 0)] = inactive
			field[vector.New4d(x, y, 1, -1)] = inactive

			field[vector.New4d(x, y, -1, 1)] = inactive
			field[vector.New4d(x, y, -1, 0)] = inactive
			field[vector.New4d(x, y, -1, -1)] = inactive
		}
	}

	for i := 0; i < 6; i++ {
		field = step4(field, false)
	}
	count := 0
	for _, v := range field {
		if v == active {
			count++
		}
	}

	return fmt.Sprint(count)
}

package main

import (
	"fmt"
	"math"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

const testData2 = `neeenesenwnwwswnenewnwwsewnenwseswesw
sewnenenenesenwsewnenwwwse`

func main() {
	input := util.ReadLines("input.txt")
	// input = strings.Split(testData, "\n")
	startTime := time.Now()
	res, field := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(field)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type Tile struct {
	Black     bool
	flipCount int

	coord complex64
}

type direction int

const (
	goEast direction = iota
	goSouthEast
	goSouthWest
	goWest
	goNorthWest
	goNorthEast
)

func parseInstructions(instructions string) []direction {
	skip := 0
	var out []direction
	for i, r := range instructions {
		if skip > 0 {
			skip--
			continue
		}
		var rNext rune
		if i+1 < len(instructions) {
			rNext = rune(instructions[i+1])
		}

		switch r {
		case 'e':
			out = append(out, goEast)
		case 'w':
			out = append(out, goWest)
		// Now we come to the tricky part
		case 'n':
			switch rNext {
			case 'e':
				out = append(out, goNorthEast)
				skip++
			case 'w':
				out = append(out, goNorthWest)
				skip++
			default:
				panic("Unknown direction")
			}

		case 's':
			switch rNext {
			case 'e':
				out = append(out, goSouthEast)
				skip++
			case 'w':
				out = append(out, goSouthWest)
				skip++
			default:
				panic("Unknown direction")
			}

		}

	}

	return out
}

var transforms = map[direction]complex64{
	goEast:      1 + 0i,
	goWest:      -1 + 0i,
	goNorthEast: 1 + 1i,
	goNorthWest: 0 + 1i,
	goSouthEast: 0 + -1i,
	goSouthWest: -1 + -1i,
}

func part1(input []string) (string, map[complex64]*Tile) {
	field := map[complex64]*Tile{}

	var startCoords complex64 = 0 + 0i
	lastCoord := startCoords

	var allInstructions [][]direction

	for _, v := range input {
		allInstructions = append(allInstructions, parseInstructions(v))
	}

	for _, instructions := range allInstructions {
		for _, dir := range instructions {
			newCoords := lastCoord + transforms[dir]
			_, exists := field[newCoords]
			if !exists {
				field[newCoords] = &Tile{coord: newCoords}
			}

			lastCoord = newCoords
		}

		targetTile := field[lastCoord]
		targetTile.Black = !targetTile.Black
		targetTile.flipCount++
		lastCoord = startCoords
	}

	blackCount := 0

	for _, t := range field {
		if t.Black {
			blackCount++
		}
	}

	return fmt.Sprint(blackCount), field
}

func fieldExtents(field map[complex64]*Tile) (int, int, int, int) {
	maxX, maxY := math.MinInt64, math.MinInt64
	minX, minY := math.MaxInt64, math.MaxInt64

	for coord := range field {
		x, y := real(coord), imag(coord)
		maxX = util.Max(int(x), maxX)
		maxY = util.Max(int(y), maxY)

		minX = util.Min(int(x), minX)
		minY = util.Min(int(y), minY)

	}

	return maxX, maxY, minX, minY
}

func part2(field map[complex64]*Tile) string {
	maxX, maxY, minX, minY := fieldExtents(field)
	changes := []complex64{}
	for i := 0; i < 100; i++ {
		for y := minY - 1; y < maxY+2; y++ {
			for x := minX - 1; x < maxX+2; x++ {
				coord := complex64(complex(float64(x), float64(y)))
				tile := field[coord]
				if tile == nil {
					tile = &Tile{coord: coord}
					field[coord] = tile

				}
				count := 0
				for _, v := range transforms {
					n, e := field[coord+v]
					if e && n.Black {
						count++
					}
				}

				if tile.Black && (count == 0 || count > 2) {
					changes = append(changes, coord)
				} else if !tile.Black && count == 2 {
					changes = append(changes, coord)
				}

			}
		}
		for _, c := range changes {
			target, exists := field[c]
			if !exists {
				field[c] = &Tile{coord: c}
				target = field[c]
			}

			target.Black = !target.Black
		}

		changes = changes[:0]

		maxX += 1
		maxY += 1
		minX -= 1
		minY -= 1
	}

	blackCount := 0

	for _, t := range field {
		if t.Black {
			blackCount++
		}
	}

	return fmt.Sprint(blackCount)
}

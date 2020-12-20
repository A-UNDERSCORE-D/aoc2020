package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

type Tile struct {
	id    int
	field [][]rune
}

func NewTile(raw []string) *Tile {
	var id int

	fmt.Sscanf(raw[0], "Tile %d:", &id)
	outField := make([][]rune, len(raw[1:]))
	for y, row := range raw[1:] {
		outField[y] = make([]rune, len(row))
		for x, r := range row {
			outField[y][x] = r
		}
	}

	return &Tile{
		id:    id,
		field: outField,
	}
}

func (t *Tile) String() string {
	out := strings.Builder{}
	out.WriteString(fmt.Sprintf("Tile %d:\n", t.id))
	for _, row := range t.field {
		for _, col := range row {
			out.WriteRune(col)
		}

		out.WriteRune('\n')
	}

	return out.String()
}

func (t *Tile) Copy() *Tile {
	outField := make([][]rune, len(t.field))
	for i, v := range t.field {
		copy(outField[i], v)
	}

	return &Tile{
		id:    t.id,
		field: outField,
	}
}

func (t *Tile) InPlaceFlipHoriz() *Tile {
	for _, row := range t.field {
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
	}
	return t
}

func (t *Tile) InPlaceFlipVert() *Tile {
	for i, j := 0, len(t.field)-1; i < j; i, j = i+1, j-1 {
		t.field[i], t.field[j] = t.field[j], t.field[i]
	}
	return t
}

func (t *Tile) AllPossibleFlipCombos() []*Tile {
	return []*Tile{
		t.Copy(),
		t.Copy().InPlaceFlipVert(),
		t.Copy().InPlaceFlipHoriz(),
		t.Copy().InPlaceFlipVert().InPlaceFlipHoriz(),
	}
}

func (t *Tile) SideMatches(side int, other *Tile) bool {
	return runeSliceEq(t.Side(side), other.Side(side))
}

func (t *Tile) AnySideMatches(other *Tile) bool {
	for _, side := range sides {
		if t.SideMatches(side, other) {
			return true
		}
	}
	return false
}

func (t *Tile) CouldMatch(other *Tile) bool {
	tF := t.AllPossibleFlipCombos()
	oF := other.AllPossibleFlipCombos()

	for _, tf := range tF {
		for _, of := range oF {
			if tf.AnySideMatches(of) {
				return true
			}
		}
	}
	return false
}

const (
	sideUp = iota
	sideDown
	sideLeft
	sideRight
)

var sides = []int{sideUp, sideDown, sideLeft, sideRight}

func (t *Tile) Side(s int) []rune {
	var out []rune
	switch s {
	case sideUp:
		out = make([]rune, len(t.field[0]))
		copy(out, t.field[0])
		return out

	case sideDown:
		out = make([]rune, len(t.field[0]))
		copy(out, t.field[len(t.field)-1])
		return out
	case sideLeft:
		out = make([]rune, len(t.field))
		for i, row := range t.field {
			out[i] = row[0]
		}
		return out
	case sideRight:
		out = make([]rune, len(t.field))
		for i, row := range t.field {
			out[i] = row[len(row)-1]
		}
		return out
	}

	return nil
}

func runeSliceEq(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func main() {
	// testStuff()
	input := strings.Split(util.ReadEntireFile("input.txt"), "\n\n")
	// input = strings.Split(fullTestData, "\n\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func (t *Tile) getAllPossibleEdges() []edge {
	PossibleSides := []edge{}
	for _, v := range sides {
		PossibleSides = append(PossibleSides, t.Side(v))
	}
	outSides := []edge{}
	for _, v := range PossibleSides {
		vCopy := make(edge, len(v))
		copy(vCopy, v)
		outSides = append(outSides, vCopy)
		vFlipped := make(edge, len(v))
		copy(vFlipped, v)

		for i, j := 0, len(vFlipped)-1; i < j; i, j = i+1, j-1 {
			vFlipped[i], vFlipped[j] = vFlipped[j], vFlipped[i]
		}
		outSides = append(outSides, vFlipped)
	}

	return outSides
}

func (t *Tile) Edges() (out []edge) {
	for _, v := range sides {
		out = append(out, t.Side(v))
	}
	return out
}

type edge = []rune

func dupCount(a, b []edge) (out int) {
	for _, x := range a {
		for _, y := range b {
			if runeSliceEq(x, y) {
				out++
			}
		}
	}

	return
}

func reversedRuneSlice(in edge) edge {
	out := make(edge, len(in))
	copy(out, in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func p1Matches(a, b edge) bool {
	return runeSliceEq(a, b) || runeSliceEq(a, reversedRuneSlice(b))
}

func anyP1Matches(a edge, b ...edge) bool {
	for _, v := range b {
		if p1Matches(a, v) {
			return true
		}
	}

	return false
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func doPart1(tiles []*Tile) int {
	out := 1
	for _, t := range tiles {
		var top, bottom, left, right bool
		t1Sides := t.Edges()
		for _, t2 := range tiles {
			if t.id == t2.id {
				continue
			}
			t2Sides := t2.Edges()

			top = anyP1Matches(t1Sides[0], t2Sides...) || top
			bottom = anyP1Matches(t1Sides[1], t2Sides...) || bottom
			left = anyP1Matches(t1Sides[2], t2Sides...) || left
			right = anyP1Matches(t1Sides[3], t2Sides...) || right
		}

		count := boolToInt(top) + boolToInt(bottom) + boolToInt(left) + boolToInt(right)
		if count == 2 {
			out *= t.id
			fmt.Println(count)
		}
	}
	return out
}

func part1(input []string) string {
	tiles := []*Tile{}
	for _, v := range input {
		tiles = append(tiles, NewTile(strings.Split(v, "\n")))
	}

	fmt.Println(doPart1(tiles))

	return ""

	tileEdges := map[int][]edge{}
	for _, tile := range tiles {
		tileEdges[tile.id] = tile.getAllPossibleEdges()
	}
	twos := map[struct{ a, b int }]struct{}{}
	for id, possibleEdges := range tileEdges {
		for id2, possibleEdges2 := range tileEdges {
			if id == id2 {
				continue
			}
			if d := dupCount(possibleEdges, possibleEdges2); d == 2 {
				_, exists1 := twos[struct{ a, b int }{id, id2}]
				_, exists2 := twos[struct{ a, b int }{id2, id}]
				if !(exists1 || exists2) {
					twos[struct{ a, b int }{id, id2}] = struct{}{}
				}
			}
		}
	}

	nums := []int{}
	for k := range twos {
		if !util.IntSliceContains(nums, k.a) {
			nums = append(nums, k.a)
		}

		if !util.IntSliceContains(nums, k.b) {
			nums = append(nums, k.b)
		}
	}

	fmt.Println(nums)

	return "stuff"
}

func part2(input []string) string {
	return "stuff2"
}

const fullTestData = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

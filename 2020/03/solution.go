package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

var testInput = strings.Split(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`, "\n")

func main() {
	input := util.ReadLines("input.txt")
	// input = testInput
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

const (
	free = iota
	tree
)

type point struct {
	x int
	y int
}

func (p point) Add(other point) point {
	return point{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p point) Sub(other point) point {
	return point{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}

type slope struct {
	slope point
	point
}

func newSlope(x, y int) *slope {
	return &slope{slope: point{x: x, y: y}}
}

func (s *slope) next() point {
	s.point = s.point.Add(s.slope)
	return s.point
}

type block struct {
	point
	content int
	hit     bool
}

func newBlock(position point, content rune) block {
	out := block{
		point: position,
	}

	switch content {
	case '.':
		out.content = free
	case '#':
		out.content = tree
	default:
		panic(fmt.Sprintf("Unknown rune %q", content))
	}

	return out
}

type field struct {
	field       map[point]block
	height      int
	width       int
	oHeight     int
	oWidth      int
	timesCopied int
}

func newField(input []string) *field {
	out := field{}
	out.parse(input)
	return &out
}

func (f *field) parse(input []string) {
	if f.field == nil {
		f.field = make(map[point]block)
	}

	for y, line := range input {
		if len(line) == 0 {
			panic("Empty line")
		}
		f.height = y + 1
		for x, r := range line {
			f.width = x + 1
			pos := point{x: x, y: y}
			f.field[pos] = newBlock(pos, r)
		}
	}

	f.oHeight = f.height
	f.oWidth = f.width
}

func (f *field) getPos(x, y int) block {
	if x < f.width && y < f.height {
		// we can access directly
		target := point{
			x: x,
			y: y,
		}
		res := f.field[target]
		if res.point != target {
			panic("Index out of bounds")
		}
		return res
	}

	f.timesCopied++
	// We need to make a new copy of our start
	for y := 0; y < f.oHeight; y++ {
		for x := 0; x < f.oWidth; x++ {
			toSet := f.getPos(x, y)
			toSet.hit = false
			newX := f.width + x
			newPos := point{
				x: newX,
				y: y,
			}

			toSet.point = newPos
			f.field[newPos] = toSet
		}
	}
	f.width += f.oWidth

	return f.getPos(x, y)
}

func (f *field) hit(pos point) {
	b := f.getPos(pos.x, pos.y)
	b.hit = true
	f.field[b.point] = b
}

func (f *field) printAtPos(pos point) {
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			p := point{x: x, y: y}
			b := f.field[p]
			if b.point != p {
				panic("asd")
			}

			r := "."
			if b.content == tree {
				r = "#"
			}

			if b.hit {
				if r == "." {
					r = "O"
				} else {
					r = "X"
				}
			}

			if p == pos {
				r = "*"
				if b.content == tree {
					r = "!"
				}
			}
			fmt.Print(r)
		}
		fmt.Println()
	}
}

func part1(input []string) string {
	f := newField(input)
	treeCount := 0

	slope := slope{
		slope: point{
			x: 3,
			y: 1,
		},
	}

	for slope.y < f.height {
		b := f.getPos(slope.x, slope.y)
		f.hit(b.point)
		// fmt.Printf("\033[2J%d, %d: Tree: %t. Total: %d\n", slope.x, slope.y, b.content == tree, treeCount)
		if b.content == tree {
			treeCount++
		}
		// f.printAtPos(b.point)

		slope.next()
		// time.Sleep(time.Millisecond * 250)
	}

	b := f.getPos(slope.x, slope.y-1)
	if b.content == tree {
		treeCount++
	}

	b = f.getPos(slope.x+3, slope.y-1)
	// fmt.Printf("\033[2J%d, %d: Tree: %t. Total: %d\n", slope.x, slope.y, b.content == tree, treeCount)
	// time.Sleep(time.Second * 3)
	return fmt.Sprint(treeCount)
}

func part2(input []string) string {
	f := newField(input)
	treeCount := 0

	// f.getPos(45, 10)
	// f.printAtPos(point{0, 0})
	slopes := []*slope{
		newSlope(1, 1),
		newSlope(3, 1),
		newSlope(5, 1),
		newSlope(7, 1),
		newSlope(1, 2),
	}

	nums := []int{}

	for _, slope := range slopes {
		// var prevPos point
		for slope.y < f.height {
			b := f.getPos(slope.x, slope.y)
			f.hit(b.point)
			// fmt.Printf("\033[2JSlope: %d; %d; %d, %d: Tree: %t. Total: %d\n", slope.slope.x, slope.slope.y, slope.x, slope.y, b.content == tree, treeCount)
			if b.content == tree {
				treeCount++
			}
			// f.printAtPos(b.point)
			// prevPos = b.point
			slope.next()
			// time.Sleep(time.Millisecond * 250)
		}

		// b := f.getPos(prevPos.x, prevPos.y)
		nums = append(nums, treeCount)
		// fmt.Printf("\033[2J%d, %d: Tree: %t. Total: %d\n", slope.x, slope.y, b.content == tree, treeCount)
		// f.printAtPos(point{x: slope.x, y: slope.y})
		treeCount = 0
	}

	total := nums[0]
	for _, n := range nums[1:] {
		total *= n
	}

	return fmt.Sprint(treeCount, nums, total)
}

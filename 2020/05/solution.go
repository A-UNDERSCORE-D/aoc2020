package main

import (
	"fmt"
	"math"
	"sort"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type boardingPass struct {
	source string
	row    int
	seat   int
}

const (
	RowMax  = 128
	SeatMax = 8
)

func makeInts(length int) (out []int) {
	for i := 0; i < length; i++ {
		out = append(out, i)
	}
	return
}

func (b *boardingPass) resolve() {
	rows := makeInts(RowMax)
	seats := makeInts(SeatMax)
	for _, chr := range b.source {
		rowCenter := int(math.Round(float64(len(rows)) / 2))
		seatCenter := int(math.Round(float64(len(seats)) / 2))
		switch chr {
		case 'F':
			rows = rows[:rowCenter]
		case 'B':
			rows = rows[rowCenter:]
		case 'L':
			seats = seats[:seatCenter]
		case 'R':
			seats = seats[seatCenter:]
		}
	}

	if len(rows) != 1 || len(seats) != 1 {
		fmt.Printf("%#v did not resolve correctly!", b)
	}
	b.row = rows[0]
	b.seat = seats[0]
}

func (b *boardingPass) id() int {
	return (b.row * 8) + b.seat
}

func part1(input []string) string {
	maxID := 0
	for _, l := range input {
		b := &boardingPass{source: l}
		b.resolve()
		maxID = util.Max(maxID, b.id())
	}
	return fmt.Sprint(maxID)
}

func part2(input []string) string {
	IDs := []int{}
	for _, l := range input {
		b := &boardingPass{source: l}
		b.resolve()
		IDs = append(IDs, b.id())
	}
	sort.Ints(IDs)
	start := IDs[0]
	for i, id := 0, start; i < len(IDs); i++ {
		if IDs[i] != id {
			return fmt.Sprint(id)
		}
		id++
	}
	return "?????"
}

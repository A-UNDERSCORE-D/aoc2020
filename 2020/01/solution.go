package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.ReadInts("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input []int) string {
	var res [2]int
	for _, outer := range input {
		for _, inner := range input {
			if outer+inner == 2020 {
				res[0], res[1] = outer, inner
				break
			}
		}
	}

	return fmt.Sprintf("%d ? %d: +: %d, *: %d", res[0], res[1], res[0]+res[1], res[0]*res[1])
}

func part2(input []int) string {
	var res [3]int
	for i, one := range input {
		for j, two := range input[i+1:] {
			for _, three := range input[j+1:] {
				if one+two+three == 2020 {
					res[0], res[1], res[2] = one, two, three
					break
				}
			}
		}
	}

	return fmt.Sprintf(
		"%d ? %d ? %d: +: %d, *: %d",
		res[0], res[1], res[2],
		res[0]+res[1]+res[2],
		res[0]*res[1]*res[2])
}

package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.GetInts(util.ReadLines("input.txt"))
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func transformNum(n int64, subj int64) int64 {
	i := n * subj
	return i % 20201227
}

func findLoopSize(targetNum int64) int {
	var transformed int64 = 7
	for i := 1; ; i++ {
		transformed = transformNum(int64(transformed), 7)
		if transformed == targetNum {
			return i + 1
		}
	}
}

func part1(input []int) string {
	card := int64(input[0])
	door := int64(input[1])
	loopSizeDoor := findLoopSize(door)

	var num int64 = 1
	for i := 0; i < loopSizeDoor; i++ {
		num = transformNum(num, card)
	}
	return fmt.Sprint(num)
}

func part2(input []int) string {
	return "This one is free :D"
}

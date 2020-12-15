package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	inputStr := util.ReadLines("input.txt")[0]
	_ = inputStr
	input := util.GetInts(strings.Split(inputStr, ","))

	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func playGame(input []int, iterations int64) int {
	LastNumberSpoken := 0
	lastSpokenMap := make(map[int]int, iterations/10)

	prevNum := -1

	for i := 1; int64(i) <= iterations; i++ {
		if i <= len(input) {
			// First turns add the original numbers
			LastNumberSpoken = input[i-1]
			lastSpokenMap[LastNumberSpoken] = i
			continue
		}

		var toSpeak int

		if res, ok := lastSpokenMap[LastNumberSpoken]; ok {
			// This has been spoken before, when?
			previousTime := i - 1
			prePreviousTime := res
			toSpeak = previousTime - prePreviousTime
		} else {
			// Not been spoken before, or only spoken once
			toSpeak = 0
		}

		if prevNum != -1 {
			// set the number we saw last loop, now that we've worked on it already
			lastSpokenMap[prevNum] = i - 1
		}

		LastNumberSpoken = toSpeak
		prevNum = toSpeak
	}
	return LastNumberSpoken
}

func part1(input []int) string {
	return strconv.FormatInt(int64(playGame(input, 2020)), 10)
}

func part2(input []int) string {
	return strconv.FormatInt(int64(playGame(input, 3e7)), 10)
}

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

func part1(input []int) string {
	lastNumberSpoken := 0
	lastSpokenMap := make(map[int][]int)

	for i := 1; i <= 2020; i++ {
		if i <= len(input) {
			// First turns add the original numbers
			lastNumberSpoken = input[i-1]
			lastSpokenMap[input[i-1]] = append(lastSpokenMap[input[i-1]], i)
			continue
		}

		var toSpeak int

		if res, ok := lastSpokenMap[lastNumberSpoken]; ok && len(res) > 1 {
			// This has been spoken before, when?
			previousTime := res[len(res)-1]
			prePreviousTime := res[len(res)-2]
			toSpeak = previousTime - prePreviousTime
		} else {
			// Not been spoken before, or only spoken once
			toSpeak = 0
		}

		lastNumberSpoken = toSpeak
		lastSpokenMap[toSpeak] = append(lastSpokenMap[toSpeak], i)
	}

	return strconv.FormatInt(int64(lastNumberSpoken), 10)
}

func part2(input []int) string {
	LastNumberSpoken := 0
	lastSpokenMap := make(map[int][]int)

	for i := 1; i <= 30000000; i++ {
		if i <= len(input) {
			// First turns add the original numbers
			// numberSpoken = append(numberSpoken, input[i-1])
			LastNumberSpoken = input[i-1]
			lastSpokenMap[input[i-1]] = append(lastSpokenMap[input[i-1]], i)
			continue
		}

		var toSpeak int

		if res, ok := lastSpokenMap[LastNumberSpoken]; ok && len(res) > 1 {
			// This has been spoken before, when?
			previousTime := res[len(res)-1]
			prePreviousTime := res[len(res)-2]
			toSpeak = previousTime - prePreviousTime
		} else {
			// Not been spoken before, or only spoken once
			toSpeak = 0
		}

		LastNumberSpoken = toSpeak
		lastSpokenMap[toSpeak] = append(lastSpokenMap[toSpeak], i)
	}

	return strconv.FormatInt(int64(LastNumberSpoken), 10)
}

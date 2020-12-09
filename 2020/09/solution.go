package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func main() {
	input := util.ReadInts("input.txt")
	// input = util.GetInts(strings.Split(testData, "\n"))
	startTime := time.Now()
	res, num := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input, num)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

const prevLength = 55

func part1(input []int) (string, int) {
outer:
	for i := prevLength; i < len(input); i++ {
		// Are we a sum of anyone?
		for j := i - prevLength; j < i+prevLength && j < len(input); j++ {
			for k := j; k < i+prevLength && k < len(input); k++ {
				if input[i] == input[j]+input[k] {
					continue outer
				}
			}
		}
		// We didnt match
		return fmt.Sprint("Weird number is: ", input[i]), input[i]
	}
	return "???", -1
}

func part2(input []int, targetNum int) string {
	var nums []int
	var sum int
	for i := 0; i < len(input); i++ {
		sum = 0
		nums = []int{}
		for j := i + 1; j < len(input); j++ {
			sum += input[j]
			nums = append(nums, input[j])
			if sum == targetNum {
				return fmt.Sprint(util.MaxOf(nums...) + util.MinOf(nums...))
			} else if sum > targetNum {
				break // Too big, try again
			}
		}
	}

	return "???"
}

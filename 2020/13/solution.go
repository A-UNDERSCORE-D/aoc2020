package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testInput = `939
1789,37,47,1889`

func main() {
	input := util.ReadLines("input.txt")
	_ = strings.Split(testInput, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input []string) string {
	curTime := util.GetInt(input[0])
	busTimes := util.GetInts(util.FilterStrSlice(strings.Split(input[1], ","), func(s string) bool { return s != "x" }))

	firstBus := -1
	atTime := -1

outer:
	for i := curTime; ; i++ {
		for _, b := range busTimes {
			if i%b == 0 {
				firstBus = b
				atTime = i
				break outer
			}
		}
	}

	return fmt.Sprint((atTime - curTime) * firstBus)
}

func part2(input []string) string {
	nums := []int{}
	for _, v := range strings.Split(input[1], ",") {
		if v == "x" {
			nums = append(nums, -1)
			continue
		}
		nums = append(nums, util.GetInt(v))
	}

	// p2 := big.NewInt(0)
	// var step int64 = 1
	// for i, bus := range nums {
	// 	if bus == -1 {
	// 		continue
	// 	}
	// 	for {
	// 		toTest := big.NewInt(0).Set(p2)
	// 		toTest.Add(toTest, big.NewInt(int64(i))).Mod(toTest, big.NewInt(int64(bus)))
	// 		if toTest.Cmp(big.NewInt(0)) == 0 {
	// 			break
	// 		}
	// 		p2.Add(p2, big.NewInt(step))
	// 	}
	// 	step *= int64(bus)
	// }

	step := 1
	var out int
	for i, bus := range nums {
		if bus == -1 {
			continue
		}

		for (out+i)%bus != 0 {
			out += step
		}
		step *= bus
	}

	return fmt.Sprint(out)
}

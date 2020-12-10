package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

const testData2 = `16
10
15
5
1
11
7
19
6
12
4`

func main() {
	input := util.ReadInts("input.txt")
	_ = util.GetInts(strings.Split(testData, "\n"))
	sort.Ints(input)
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input []int) string {
	realData := make([]int, 0, len(input)+1)
	realData = append(realData, 0)
	realData = append(realData, input...)
	realData = append(realData, realData[len(realData)-1]+3)

	ones := 0
	threes := 0
	last := realData[0]
	for _, adapter := range realData[1:] {
		switch diff := adapter - last; diff {
		case 1:
			ones++
		case 3:
			threes++
		default:
			fmt.Printf("Unexpected difference %d\n", diff)
		}
		last = adapter
	}
	return strconv.FormatInt(int64(ones*threes), 10)
}

func part2(input []int) string {
	cache := map[int]int{}
	var recurse func(int) int
	realData := []int{0}
	realData = append(realData, input...)

	recurse = func(idx int) (out int) {
		if idx == len(realData)-1 {
			return 1
		}
		if res, ok := cache[idx]; ok {
			return res
		}

		for i := idx + 1; i < len(realData); i++ {
			if realData[i]-realData[idx] <= 3 {
				out += recurse(i)
			} else {
				break
			}
		}

		cache[idx] = out
		return
	}

	return strconv.FormatInt(int64(recurse(0)), 10)
}

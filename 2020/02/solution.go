package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

var testData = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func main() {
	input := util.ReadLines("input.txt")
	// input = strings.Split(testData, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type password struct {
	password string
	checkChr rune
	minCount int
	maxCount int
	counts   map[rune]int
}

func (p *password) String() string {
	return fmt.Sprintf("password{%s, %s, max: %d, min: %d}", p.password, string(p.checkChr), p.maxCount, p.minCount)
}

func (p *password) countChrs() {
	p.counts = make(map[rune]int)
	for _, r := range p.password {
		p.counts[r]++
	}
}

func (p *password) isValidPt1() bool {
	p.countChrs()
	return !(p.counts[p.checkChr] > p.maxCount || p.counts[p.checkChr] < p.minCount)
}

func (p *password) isValidPt2() bool {
	if len(p.password) < p.maxCount {
		return false
	}

	one, two := rune(p.password[p.minCount-1]), rune(p.password[p.maxCount-1])
	if (one == p.checkChr || two == p.checkChr) && two != one {
		return true
	}
	return false
}

func parsePasswd(raw string) *password {
	split := strings.Fields(raw)
	minMax := strings.Split(split[0], "-")
	return &password{
		password: split[2],
		checkChr: rune(split[1][0]),
		minCount: util.GetInt(minMax[0]),
		maxCount: util.GetInt(minMax[1]),
	}
}

func part1(input []string) string {
	count := 0
	for _, l := range input {
		if parsePasswd(l).isValidPt1() {
			count++
		}
	}
	return fmt.Sprint(count)
}

func part2(input []string) string {
	count := 0
	for _, l := range input {
		if p := parsePasswd(l); p.isValidPt2() {
			count++
		}
	}
	return fmt.Sprint(count)
}

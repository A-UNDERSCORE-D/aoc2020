package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

const testp2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func main() {
	input := strings.Split(util.ReadEntireFile("input.txt"), "\n\n")
	// input = strings.Split(testp2, "\n\n")
	startTime := time.Now()
	rules := parseRules(strings.Split(input[0], "\n"))
	fmt.Println("Rules parsed. Took:", time.Since(startTime))

	startTime = time.Now()
	res, validTickets := part1(rules, input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part2(input, validTickets, rules)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type Rule struct {
	name   string
	ranges [2][2]int
}

func (r *Rule) numIsValid(num int) bool {
	// I know. Its here for debugging
	if (num >= r.ranges[0][0] && num <= r.ranges[0][1]) || (num >= r.ranges[1][0] && num <= r.ranges[1][1]) {
		return true
	}
	return false
}

func (r *Rule) numsAreValid(nums ...int) bool {
	for _, v := range nums {
		if !r.numIsValid(v) {
			return false
		}
	}
	return true
}

func (r *Rule) isValid(ticket []int) (bool, int) {
	for _, v := range ticket {
		if !r.numIsValid(v) {
			return false, v
		}
	}
	return true, -1
}

func parseRules(ruleStrs []string) []*Rule {
	var rules []*Rule
	for _, v := range ruleStrs {
		split := strings.Split(v, ": ")
		name := split[0]
		rangeSplit := strings.Split(split[1], " or ")
		ranges := [2][2]int{}

		for i, v := range rangeSplit {
			minMax := util.GetInts(strings.Split(v, "-"))
			ranges[i] = [2]int{minMax[0], minMax[1]}
		}

		rules = append(rules, &Rule{
			name:   name,
			ranges: ranges,
		})

	}

	return rules
}

func part1(rules []*Rule, input []string) (string, [][]int) {
	otherTicketsStr := strings.Split(input[2], "\n")[1:]
	var otherTickets [][]int
	for _, v := range otherTicketsStr {
		otherTickets = append(otherTickets, util.GetInts(strings.Split(v, ",")))
	}

	var invalidValues []int
	var validTickets [][]int

	for _, ticket := range otherTickets {
		startLen := len(invalidValues)
	ticketLoop:
		for _, value := range ticket {
			for _, rule := range rules {
				if rule.numIsValid(value) {
					continue ticketLoop
				}
			}
			invalidValues = append(invalidValues, value)
		}

		if len(invalidValues) == startLen {
			// Ticket had no invalid spots
			validTickets = append(validTickets, ticket)
		}
	}

	return strconv.FormatInt(int64(util.Sum(invalidValues...)), 10), validTickets
}

func part2(input []string, tickets [][]int, rules []*Rule) string {
	ourTicket := util.GetInts(strings.Split(strings.Split(input[1], "\n")[1], ","))
	var transposedIdx [][]int

	for i := 0; i < len(tickets[0]); i++ {
		var nums []int
		for _, v := range tickets {
			nums = append(nums, v[i])
		}
		transposedIdx = append(transposedIdx, nums)
	}

	candidateMap := make(map[*Rule][]int)
	for _, rule := range rules {
		for i, col := range transposedIdx {
			if rule.numsAreValid(col...) {
				candidateMap[rule] = append(candidateMap[rule], i)
			}
		}
	}

	needWork := true
	// Okay we have all possible candidates for every rule, now lets work backwards from the rules that only
	// have one possible candidate, and remove it from all others until every rule only has one candidate
	for needWork {
		for _, rule := range rules {
			candidates := candidateMap[rule]
			if len(candidates) > 1 {
				continue
			} else if len(candidates) == 0 {
				panic("No candidates found")
			}

			correctOne := candidates[0]
			// newCMap := make(map[rule][]int, len(candidateMap))
			for k, otherCandidates := range candidateMap {
				if len(otherCandidates) == 1 || k == rule {
					continue
				}
				for i, c := range otherCandidates {
					if c == correctOne {
						candidateMap[k] = append(otherCandidates[:i], otherCandidates[i+1:]...)
						break
					}
				}
			}
		}

		needWork = false
		for _, r := range rules {
			// Are we done?
			if len(candidateMap[r]) > 1 {
				needWork = true
				break
			}
		}
	}

	total := -1

	for k, v := range candidateMap {
		if strings.HasPrefix(k.name, "departure") {
			target := v[0]
			if total == -1 {
				total = ourTicket[target]
			} else {
				total *= ourTicket[target]
			}
		}
	}

	return strconv.FormatInt(int64(total), 10)
}

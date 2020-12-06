package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"awesome-dragon.science/go/adventofcode2020/util/set"
)

const testData = `abc

a
b
c

ab
ac

a
a
a
a

b`

func main() {
	f, err := os.Create("profile")
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

	input := strings.Split(util.ReadEntireFile("input.txt"), "\n\n")
	// input = strings.Split(testData, "\n\n")

	input2 := make([]string, len(input))
	copy(input2, input)
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part1v2(input)
	fmt.Println("Part 1v2:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part2(input2)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part2v2(input2)
	fmt.Println("Part 2v2:", res, "Took:", time.Since(startTime))
}

func part1(input []string) string {
	totalAnswers := 0
	for _, group := range util.ReplaceAllSlice(input, "\n", "") {
		seen := make(map[rune]struct{})
		for _, question := range group {
			if _, ok := seen[question]; !ok {
				seen[question] = struct{}{}
				totalAnswers++
			}
		}
	}

	return fmt.Sprint(totalAnswers)
}

func part1v2(input []string) string {
	totalAnswers := 0
	for _, group := range util.ReplaceAllSlice(input, "\n", "") {
		groupSet := set.NewSet()
		for _, question := range group {
			groupSet.Insert(question)
		}

		totalAnswers += groupSet.Length()

	}
	return fmt.Sprint(totalAnswers)
}

func part2(input []string) string {
	totalAnswers := 0
	for _, group := range input {
		groupAnswers := make(map[rune]int)
		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, question := range person {
				groupAnswers[question]++
			}
		}

		groupTotal := 0
		for _, answerCount := range groupAnswers {
			if answerCount == len(people) {
				groupTotal++
			}
		}

		totalAnswers += groupTotal
	}

	return fmt.Sprint(totalAnswers)
}

func part2v2(input []string) string {
	total := 0
	for _, group := range input {
		answers := []*set.Set{}
		people := strings.Split(group, "\n")
		for _, person := range people {
			personSet := set.NewSet()
			for _, q := range person {
				personSet.Insert(q)
			}

			answers = append(answers, personSet)
		}

		var groupTotal *set.Set
		for _, set := range answers {
			if groupTotal == nil {
				groupTotal = set
				continue
			}
			groupTotal = groupTotal.Intersect(set)
		}

		total += groupTotal.Length()
	}
	return fmt.Sprint(total)
}

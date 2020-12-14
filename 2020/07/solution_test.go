package main

import (
	"testing"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func BenchmarkPart1(b *testing.B) {
	input := util.ReadLines("input.txt")
	bags := parseBagsGraph(input)
	for i := 0; i < b.N; i++ {
		part1(bags)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.ReadLines("input.txt")
	bags := parseBagsGraph(input)
	for i := 0; i < b.N; i++ {
		part2(bags)
	}
}

func BenchmarkParse(b *testing.B) {
	input := util.ReadLines("input.txt")
	for i := 0; i < b.N; i++ {
		parseBagsGraph(input)
	}
}

package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func main() {
	input := util.ReadLines("input.txt")
	_ = strings.Split(testData, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func applyMask(mask string, input uint64) uint64 {
	mask = strings.TrimPrefix(mask, "mask = ")
	// fmt.Printf("value  = %064b (decimal %[1]d)\n", input)
	// fmt.Printf("mask   = %064s\n", mask)
	for i, v := range mask {
		switch v {
		case 'X':
		case '0':

			input &= ^(1 << (35 - i))
		case '1':
			input |= (1 << (35 - i))
		}
	}
	// fmt.Printf("result = %064b (decimal %[1]d)\n", input)
	return input
}

func part1(input []string) string {
	memory := make(map[int]uint64)
	currentMask := ""
	for _, v := range input {
		if strings.HasPrefix(v, "mask = ") {
			currentMask = v
			continue
		}

		memIdx := 0
		var toSet uint64 = 0
		if n, err := fmt.Sscanf(v, "mem[%d] = %d", &memIdx, &toSet); n != 2 || err != nil {
			panic(fmt.Sprint(n, err))
		}

		memory[memIdx] = applyMask(currentMask, toSet)
	}

	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	return fmt.Sprint(total)
}

func maskPermut(mask string) []maskWithFloats {
	mask = strings.TrimPrefix(mask, "mask = ")
	var out []maskWithFloats
	num := strings.Count(mask, "X")
	var realNum uint64 = 0xFFFFFFFFFFFFFFFF >> (64 - num)
	printFMask := fmt.Sprintf("%%0%db", num)
	nextMask := strings.Builder{}
	nextMask.Grow(len(mask))
	for i := 0; uint64(i) <= realNum; i++ {
		bitCount := 0
		bits := fmt.Sprintf(printFMask, i)
		floats := [][2]int{}
		for i, chr := range mask {
			switch chr {
			case 'X':
				nextMask.WriteByte(bits[bitCount])
				floats = append(floats, [2]int{35 - i, int(bits[bitCount]) - 48})
				bitCount++
			default:
				nextMask.WriteRune(chr)
			}
		}

		out = append(out, maskWithFloats{mask: nextMask.String(), floats: floats})
		nextMask.Reset()
	}

	return out
}

func setBit(num uint64, target, bit int) uint64 {
	switch bit {
	case 0:
		num &= ^(1 << target)
	case 1:
		num |= 1 << target
	default:
		panic("Tried to set a bit that is neither 0 nor 1")
	}
	return num
}

func applyMaskP2(mask maskWithFloats, input uint64) uint64 {
	mask.mask = strings.TrimPrefix(mask.mask, "mask = ")
	for i, v := range mask.mask {
		switch v {
		case '1':
			input |= (1 << (35 - i))
		}
		for _, v := range mask.floats {
			if v[0] == i {
				input = setBit(input, v[0], v[1])
			}
		}
	}
	return input
}

type maskWithFloats struct {
	mask   string
	floats [][2]int
}

func part2(input []string) string {
	memory := make(map[uint64]uint64)
	currentMask := ""
	currentMasks := []maskWithFloats{}
	for _, v := range input {
		if strings.HasPrefix(v, "mask = ") {
			currentMask = v
			currentMasks = maskPermut(currentMask)
			continue
		}

		var memIdx uint64 = 0
		var toSet uint64 = 0
		if n, err := fmt.Sscanf(v, "mem[%d] = %d", &memIdx, &toSet); n != 2 || err != nil {
			panic(fmt.Sprint(n, err))
		}

		for _, mask := range currentMasks {
			target := applyMaskP2(mask, memIdx)
			memory[target] = toSet
		}

	}

	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	return fmt.Sprint(total)
}

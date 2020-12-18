package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`

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

type opcode rune

// const (
// 	add = '+'
// 	mul = '*'
// )

// type data struct {
// 	num int
// }

// type op struct {
// 	op       opcode
// 	constant int
// 	left     *data
// }

// func parseToOps(input string) []*op {
// 	nextOp := &op{}
// 	for _, r := range input {
// 		switch r {
// 		case ' ':
// 			continue
// 		case '+':
// 			nextOp.op = add
// 		case '*':
// 			nextOp.op = mul

// 		case '(', ')':
// 			panic("for now not implemented")
// 		default:
// 			// Must be a number
// 		}
// 	}
// }

func doLine1(line string, inParen bool) (int, int) {
	leftNum := 0
	rightNum := 0
	var op func(a, b int) int
	seenLeft := false
	seenRight := false
	skipTo := 0
	for i, r := range line {
		if skipTo > 0 {
			skipTo--
			continue
		}
		switch r {
		case ' ':
			continue

		case '+':
			op = func(a, b int) int { return a + b }
		case '*':
			op = func(a, b int) int { return a * b }
		case '(':
			// We're entering a paren, we need to resolve it before we can resolve the rest of the stuff
			num, skip := doLine1(line[i+1:], true)
			skipTo = skip
			if seenLeft {
				rightNum = num
				seenRight = true
			} else {
				leftNum = num
				seenLeft = true
			}
			// skipTo = skipTo - i // skip forward the difference between the end of the last and now
		case ')':
			if !inParen {
				panic("Invalid syntax")
			}

			// Okay we found our end, return out here with whatever is in leftNum
			return leftNum, i + 1
		default:
			num := util.GetInt(string(r))

			if seenLeft {
				rightNum = num
				seenRight = true
			} else {
				leftNum = num
				seenLeft = true
			}
		}

		if seenLeft && seenRight {
			// we have both left and right, do the op and continue
			num := op(leftNum, rightNum)
			leftNum = num
			seenRight = false
			rightNum = 0
		}
	}

	return leftNum, 0
}

func part1(input []string) string {
	sum := 0
	for _, v := range input {
		num, _ := doLine1(v, false)
		sum += num
		fmt.Println(num, "\n---")
	}

	return fmt.Sprint(sum)
}

func nextOp(line string) (rune, int) {
	idx := strings.IndexAny(line, "()+*")
	if idx == -1 {
		return 0x0, 0
	}
	return rune(line[idx]), idx
}

var numRe = regexp.MustCompile(`\d+`)

func findClosing(line string) int {
	stack := 0
	for i, v := range line {
		switch v {
		case '(':
			stack++
		case ')':
			stack--
			if stack == 0 {
				return i
			}
		}
	}
	return -1
}

func nThIndexOf(line string, target rune, n int) int {
	out := 0
	for i := 0; i < n; i++ {
		if out == -1 {
			return out
		}
		toCheck := line[out:]
		idx := strings.IndexRune(toCheck, target)
		if idx == -1 {
			return -1
		}
		out = idx + out + 1
	}

	return out - 1
}

func doLine4(line string) string {
	if !strings.ContainsAny(line, "()*+") {
		return line
	}
	leftNum := 0
	rightNum := 0
	var op func(a, b int) int
	seenLeft := false
	seenRight := false
	skipTo := 0
	for i, r := range line {
		if skipTo > 0 {
			skipTo--
			continue
		}
		switch r {
		case ' ', ')':
			continue
		case '+':
			op = func(a, b int) int { return a + b }
		case '*':
			op = func(a, b int) int { return a * b }

			nextOp, _ := nextOp(line[i+1:])
			if nextOp == '+' {
				// Add gets a higher priority, so resolve whatever that number is first
				startNext := i + 2
				endNext := 0
				lineStartNext := line[startNext:]
				if line[startNext] == '(' {
					// we need to find the end of this and resolve it out
					endNext = findClosing(line[startNext:]) + 1
				} else {
					// Its a simpler expression, just add the numbers and replace them
					// * 354 + 4654654
					//   ^           ^
					//   |           L we want endNext to be here
					//   L startnext is here
					endNext = nThIndexOf(lineStartNext, ' ', 3)
					if endNext < 0 {
						endNext = len(lineStartNext)
					}
				}

				toReplace := lineStartNext[:endNext]
				if strings.Contains(toReplace, "(") {
					endNext = findClosing(line[startNext:]) + 1
					toReplace = lineStartNext[:endNext]
				}

				newExpr := doLine4(toReplace)
				newLine := strings.Replace(line, toReplace, newExpr, 1)
				return doLine4(newLine)
			}

		case '(':
			closingIdx := findClosing(line)
			toReplace := doLine4(line[i+1 : closingIdx])
			line = strings.Replace(line, line[i:closingIdx+1], toReplace, 1)
			return doLine4(line)

		default:
			// Must be a number
			realNum := numRe.FindString(line[i:])
			if len(realNum) == 0 {
				panic("Empty number?")
			}
			skipTo = len(realNum) - 1

			num := util.GetInt(realNum)

			if seenLeft {
				rightNum = num
				seenRight = true
			} else {
				leftNum = num
				seenLeft = true
			}
		}

		if seenLeft && seenRight {
			num := op(leftNum, rightNum)
			toReplace := line[:i+1+skipTo]
			newLine := strings.Replace(line, toReplace, fmt.Sprint(num), 1)
			return doLine4(newLine)
		}
	}

	return fmt.Sprint(leftNum)
}

func part2(input []string) string {
	total := 0
	for _, v := range input {
		total += util.GetInt(doLine4(v))
	}

	return fmt.Sprint(total)
}

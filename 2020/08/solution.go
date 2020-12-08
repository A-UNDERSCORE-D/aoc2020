package main

import (
	"errors"
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func main() {
	input := util.ReadLines("input.txt")
	// input = strings.Split(testInput, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type opcode int

const (
	unknown opcode = iota
	acc
	jmp
	nop
)

var opLut = map[string]opcode{
	"acc": acc,
	"jmp": jmp,
	"nop": nop,
}

type HHGC struct {
	accumulator int
	source      []string
	compiled    []operation
}

type operation struct {
	opcode opcode
	arg    int
}

func parseHHGCCode(input []string) *HHGC {
	out := &HHGC{}
	for _, line := range input {
		if line == "" {
			continue
		}
		out.source = append(out.source, line)
		var (
			op  string
			arg int
		)
		if n, err := fmt.Sscanf(line, "%s %d", &op, &arg); err != nil || n != 2 {
			panic(fmt.Sprint("Incorrect number of values read ", err, n))
		}

		out.compiled = append(out.compiled, operation{opcode: opLut[op], arg: arg})
	}

	return out
}

var errInterrupted = errors.New("interrupted by line handler")

func (h *HHGC) Run(lineHandler func(line int) bool) error {
	for i := 0; i < len(h.compiled); i++ {
		if lineHandler != nil && !lineHandler(i) {
			return errInterrupted
		}
		op := h.compiled[i]
		switch op.opcode {
		case jmp:
			i += op.arg
			i--

		case acc:
			h.accumulator += op.arg

		case nop:
			continue
		case unknown:
			fallthrough
		default:
			panic("Unknown op!")
		}
	}
	return nil
}

func part1(input []string) string {
	vm := parseHHGCCode(input)
	seenLines := map[int]struct{}{}
	err := vm.Run(func(line int) bool {
		if _, ok := seenLines[line]; ok {
			return false
		}
		seenLines[line] = struct{}{}
		return true
	})

	if err != nil && err != errInterrupted {
		panic(err)
	}
	return fmt.Sprint(vm.accumulator)
}

func nopToJmp(op operation) operation {
	if op.opcode == nop {
		return operation{
			opcode: jmp,
			arg:    op.arg,
		}
	}
	return operation{
		opcode: nop,
		arg:    op.arg,
	}
}

func part2(input []string) string {
	vm := parseHHGCCode(input)
	lastModified := -1

	for i := 0; i < len(vm.compiled); i++ {
		vm.accumulator = 0
		ins := vm.compiled[i]

		seenLines := map[int]struct{}{}
		err := vm.Run(func(line int) bool {
			if _, ok := seenLines[line]; ok {
				return false
			}
			seenLines[line] = struct{}{}
			return true
		})

		if err == errInterrupted {
			// we looped, dont do that
			if lastModified != -1 {
				vm.compiled[lastModified] = nopToJmp(vm.compiled[lastModified])
				lastModified = -1
			}
		}

		if ins.opcode == nop || ins.opcode == jmp {
			lastModified = i
			vm.compiled[i] = nopToJmp(vm.compiled[i])
		}

		if err == nil {
			break
		}

		if err != nil && err != errInterrupted {
			panic(err)
		}
	}
	return fmt.Sprint(vm.accumulator)
}

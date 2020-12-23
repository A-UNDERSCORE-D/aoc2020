package main

import (
	"container/ring"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.DecimalExplode(util.GetInt(util.ReadLines("input.txt")[0]))
	// input = util.DecimalExplode(389125467)
	debug.SetGCPercent(-1)
	startTime := time.Now()
	res := part1(util.CopyIntSlice(input))
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(util.CopyIntSlice(input))
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type Cups []int

func (c *Cups) indexLoop(idx int) int {
	return (*c)[idx%len(*c)]
}

func (c *Cups) popIdx(idx int) int {
	normalizedIdx := idx % len(*c)
	toRet := (*c)[normalizedIdx]
	(*c) = append((*c)[:normalizedIdx], (*c)[normalizedIdx+1:]...)
	// (*c) = (*c)[:len((*c))-1]
	return toRet
}

func (c *Cups) labelIdx(idx int) string {
	out := strings.Builder{}
	for i, v := range *c {
		if idx != -1 && i == idx%len(*c) {
			out.WriteString(fmt.Sprintf("(%d)", v))
		} else {
			out.WriteString(fmt.Sprint(v))
		}

		out.WriteString(" ")
	}

	return out.String()
}

func ringContains(head *ring.Ring, value int) bool {
	current := head
	for i := 0; i < current.Len(); i++ {
		if current.Value.(int) == value {
			return true
		}
		current = current.Next()
	}
	return false
}

func ringMax(r *ring.Ring) int {
	v := 0
	r.Do(func(i interface{}) {
		num, ok := i.(int)
		if ok {
			v = util.Max(v, num)
		}
	})
	return v
}

func ringWithValue(head *ring.Ring, value int) *ring.Ring {
	current := head
	for i := 0; i < head.Len(); i++ {
		if current.Value.(int) == value {
			return current
		}
		current = current.Next()
	}
	return nil
}

func part1(input []int) string {
	cupsHead := ring.New(len(input))
	for _, v := range input {
		cupsHead.Value = v
		cupsHead = cupsHead.Next()
	}
	origHead := cupsHead
	_ = origHead

	for i := 0; i < 100; i++ {
		picked := cupsHead.Unlink(3)
		dst := cupsHead.Value.(int) - 1
		if dst == 0 {
			dst = ringMax(cupsHead)
		}

		for ringContains(picked, dst) {
			dst--
			if dst == 0 {
				dst = ringMax(cupsHead)
			}
		}

		dstRing := ringWithValue(cupsHead, dst)
		dstRing.Link(picked)
		cupsHead = cupsHead.Next()
	}

	out := strings.Builder{}

	ringWithValue(cupsHead, 1).Next().Do(func(i interface{}) {
		if i.(int) == 1 {
			return
		}
		out.WriteString(fmt.Sprint(i))
	})

	return out.String()
}

func part2(input []int) string {
	cupsHead := ring.New(1000000)
	numsLut := make(map[int]*ring.Ring)
	for _, v := range input {
		cupsHead.Value = v
		numsLut[v] = cupsHead
		cupsHead = cupsHead.Next()
	}

	for i := ringMax(cupsHead) + 1; i <= 1000000; i++ {
		cupsHead.Value = i
		numsLut[i] = cupsHead
		cupsHead = cupsHead.Next()
	}
	origHead := cupsHead
	_ = origHead

	for i := 0; i < 10000000; i++ {
		picked := cupsHead.Unlink(3)
		dst := cupsHead.Value.(int) - 1
		if dst == 0 {
			dst = ringMax(cupsHead)
		}

		for ringContains(picked, dst) {
			dst--
			if dst == 0 {
				dst = ringMax(cupsHead)
			}
		}

		dstRing, ok := numsLut[dst]
		if !ok {
			panic("Unknown number@")
		}
		dstRing.Link(picked)
		cupsHead = cupsHead.Next()
	}

	one := numsLut[1]
	a := one.Next()
	b := a.Next()

	return fmt.Sprint(a.Value.(int) * b.Value.(int))
}

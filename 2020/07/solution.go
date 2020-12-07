package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const test = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

const test2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func main() {
	input := util.ReadLines("input.txt")
	// input = strings.Split(test, "\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type unresolvedBag struct {
	colour   string
	resolved bool
	count    int
}

type Bag struct {
	colour             string
	unresolvedContains []unresolvedBag
	canContain         []struct {
		*Bag
		count int
	}
}

func (b *Bag) couldContain(name string) bool {
	for _, bag := range b.canContain {
		if bag.colour == name {
			return true
		}
	}

	for _, bag := range b.canContain {
		if bag.couldContain(name) {
			return true
		}
	}

	return false
}

func (b *Bag) maxContents() int {
	count := 0
	for _, bag := range b.canContain {
		count += bag.count
		count += bag.maxContents() * bag.count
	}

	return count
}

func parseBagsGraph(input []string) map[string]*Bag {
	outerBags := make(map[string]*Bag)

	for _, v := range input {
		split := strings.Split(v, " contain ")
		name := strings.TrimSuffix(split[0], " bags")
		contents := strings.Split(split[1][:len(split[1])-1], ", ")
		var unresolvedBags []unresolvedBag
		for _, v := range contents {
			if v == "no other bags" {
				continue
			}

			v = strings.TrimSuffix(v, "bag")
			v = strings.TrimSuffix(v, "bags")

			count, name := v[0:1], v[2:len(v)-1]
			unresolvedBags = append(unresolvedBags, unresolvedBag{colour: name, count: util.GetInt(count)})
		}
		outerBags[name] = &Bag{
			colour:             name,
			unresolvedContains: unresolvedBags,
		}

	}

	// Now, to resolve all the unresolveds
	for _, outerBag := range outerBags {
		if outerBag.unresolvedContains == nil {
			// Cant contain any others, skip it
			continue
		}

		for _, unresolved := range outerBag.unresolvedContains {
			if unresolved.resolved {
				fmt.Println("Attempted to resolve already resolved bag: ", unresolved)
				continue
			}

			bag, ok := outerBags[unresolved.colour]
			if !ok {
				panic(fmt.Sprintf("Unresolvable bag %q!", unresolved.colour))
			}
			outerBag.canContain = append(outerBag.canContain, struct {
				*Bag
				count int
			}{bag, unresolved.count})
			unresolved.resolved = true
		}
	}

	return outerBags
}

const targetBag = "shiny gold"

func part1(input []string) string {
	bags := parseBagsGraph(input)
	count := 0
	for _, bag := range bags {
		if bag.couldContain(targetBag) {
			count++
		}
	}

	return fmt.Sprint(count)
}

func part2(input []string) string {
	bags := parseBagsGraph(input)
	return fmt.Sprint(bags[targetBag].maxContents())
}

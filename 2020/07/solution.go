package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	bags := parseBagsGraph(input)
	fmt.Println("Bags parsed, took: ", time.Since(startTime))
	startTime = time.Now()
	res := part1(bags)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(bags)
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

func part1(bags map[string]*Bag) string {
	count := 0
	for _, bag := range bags {
		if bag.couldContain(targetBag) {
			count++
		}
	}

	return fmt.Sprint(count)
}

func part2(bags map[string]*Bag) string {
	return fmt.Sprint(bags[targetBag].maxContents())
}

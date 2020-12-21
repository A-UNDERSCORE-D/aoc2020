package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"awesome-dragon.science/go/adventofcode2020/util/set"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res, ings := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(ings)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type IngredientsList struct {
	ingredients []string
	allergens   []string
}

func (i *IngredientsList) String() string {
	return fmt.Sprintf("%s with allergens %s.", strings.Join(i.ingredients, ", "), strings.Join(i.allergens, ", "))
}

func NewIngredientsList(line string) *IngredientsList {
	out := new(IngredientsList)
	split := strings.Split(line, " (")
	out.ingredients = strings.Split(split[0], " ")
	if len(split) == 2 {
		clean := strings.TrimPrefix(split[1], "contains ")
		clean = clean[:len(clean)-1]
		out.allergens = strings.Split(clean, ", ")
	}

	return out
}

func makeSet(data []string) *set.StringSet {
	out := set.NewStringSet()
	for _, v := range data {
		out.Insert(v)
	}
	return out
}

func part1(input []string) (string, map[string]*set.StringSet) {
	var ingredients []*IngredientsList
	for _, v := range input {
		ingredients = append(ingredients, NewIngredientsList(v))
	}
	// ingredientsStr := set.NewSet()
	AllPossibleAllergens := map[string]*set.StringSet{}
	AllIngredients := set.NewStringSet()
	for _, v := range ingredients {
		AllIngredients = AllIngredients.Union(makeSet(v.ingredients))
		for _, a := range v.allergens {
			if AllPossibleAllergens[a] == nil {
				AllPossibleAllergens[a] = set.NewStringSet()
				// First time, add everything
				for _, i := range v.ingredients {
					AllPossibleAllergens[a].Insert(i)
				}
			} else {
				// We've seen this allergen before, now just do everything that exists in both
				AllPossibleAllergens[a] = AllPossibleAllergens[a].Intersect(makeSet(v.ingredients))
			}
		}
	}
	// Now, remove any confirmed allergen that appears in others
	work := true
	for work {

		for name, i := range AllPossibleAllergens {
			if i.Length() != 1 {
				continue // cant resolve when there are many
			}

			for other, otherI := range AllPossibleAllergens {
				if other == name {
					continue
				}

				for _, v := range i.Values() {
					otherI.Remove(v)
				}
			}
		}

		work = false
		for _, a := range AllPossibleAllergens {
			if a.Length() > 1 {
				work = true
				continue
			}
		}
	}

	allAllergensAnyType := set.NewStringSet()
	for _, v := range AllPossibleAllergens {
		allAllergensAnyType = allAllergensAnyType.Union(v)
	}

	NonAllergens := allAllergensAnyType.Difference(AllIngredients)
	count := 0
	for _, ing := range ingredients {
		for _, na := range NonAllergens.Values() {
			if util.StringSliceContains(ing.ingredients, na) {
				count++
			}
		}
	}

	return fmt.Sprint(count), AllPossibleAllergens
}

func part2(input map[string]*set.StringSet) string {
	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	allergens := []string{}
	for _, k := range keys {
		allergens = append(allergens, input[k].Values()[0])
	}

	return strings.Join(allergens, ",")
}

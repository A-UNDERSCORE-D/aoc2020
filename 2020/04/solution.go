package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func main() {
	input := util.ReadEntireFile("input.txt")
	// input = testData

	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func parseData(input string) []map[string]string {
	split := strings.Split(input, "\n")
	var out []map[string]string
	next := make(map[string]string)
	for _, v := range split {
		if len(v) == 0 {
			out = append(out, next)
			next = make(map[string]string)
			continue
		}

		for _, pair := range strings.Split(v, " ") {
			s := strings.Split(pair, ":")
			next[s[0]] = s[1]
		}
	}
	out = append(out, next)

	return out
}

func mapHasAllFields(m map[string]string, fields ...string) []string {
	missing := []string{}
	for _, field := range fields {
		if _, ok := m[field]; !ok {
			missing = append(missing, field)
		}
	}

	if len(missing) > 0 {
		return missing
	}
	return nil
}

var (
	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	validEyes      = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
)

func digitsAtLeast(in string, size int) (int, bool) {
	if len(in) != size {
		return -1, false
	}

	num, err := strconv.Atoi(in)
	if err != nil {
		return -1, false
	}

	return num, true
}

var validFuncs = map[string]func(string) bool{
	"byr": func(s string) bool {
		num, ok := digitsAtLeast(s, 4)
		if !ok {
			return false
		}
		return num >= 1920 && num <= 2002
	},
	"iyr": func(s string) bool {
		num, ok := digitsAtLeast(s, 4)
		if !ok {
			return false
		}
		return num >= 2010 && num <= 2020
	},
	"eyr": func(s string) bool {
		num, ok := digitsAtLeast(s, 4)
		if !ok {
			return false
		}
		return num >= 2020 && num <= 2030
	},
	"hgt": func(s string) bool {
		cm := strings.HasSuffix(s, "cm")
		if !cm && !strings.HasSuffix(s, "in") {
			return false
		}
		num, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false
		}

		return (cm && num >= 150 && num <= 193) || (!cm && num >= 59 && num <= 76)
	},
	"hcl": func(s string) bool {
		if !strings.HasPrefix(s, "#") {
			return false
		}
		for _, v := range strings.ToLower(s[1:]) {
			if !strings.ContainsAny(string(v), "abcdef0123456789") {
				return false
			}
		}
		return true
	},
	"ecl": func(s string) bool {
		return util.StringSliceContains(validEyes, s)
	},
	"pid": func(s string) bool {
		num, ok := digitsAtLeast(s, 9)
		if !ok {
			return false
		}

		return s == fmt.Sprintf("%09d", num)
	},
}

func part1(input string) string {
	parsed := parseData(input)
	validCount := 0
	for _, v := range parsed {
		if res := mapHasAllFields(v, requiredFields...); res == nil || len(res) == 1 && res[0] == "cid" {
			validCount++
		}
	}

	return fmt.Sprint(validCount)
}

func part2Valid(id map[string]string) bool {
	for _, field := range requiredFields[:len(requiredFields)-1] {
		if !validFuncs[field](id[field]) {
			return false
		}
	}
	return true
}

func part2(input string) string {
	parsed := parseData(input)
	p1Valid := []map[string]string{}
	for _, v := range parsed {
		if res := mapHasAllFields(v, requiredFields...); res == nil || len(res) == 1 && res[0] == "cid" {
			p1Valid = append(p1Valid, v)
		}
	}

	validCount := 0

	for _, id := range p1Valid {
		if part2Valid(id) {
			validCount++
		}
	}

	return fmt.Sprint(validCount)
}

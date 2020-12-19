package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

const testData = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

const testData2 = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

func main() {
	input := strings.Split(util.ReadEntireFile("input.txt"), "\n\n")
	// input = strings.Split(yitz, "\n\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type Rule struct {
	isLiteral  bool
	contentRaw string
	ruleNum    string
}

func NewRule(str string) *Rule {
	strSplit := strings.SplitN(str, ": ", 2)
	str = strSplit[1]
	if str[0] == '"' {
		return &Rule{
			ruleNum:    strSplit[0],
			contentRaw: str[1 : len(str)-1],
			isLiteral:  true,
		}
	}
	return &Rule{
		ruleNum:    strSplit[0],
		contentRaw: str,
	}
}

func (r *Rule) resolve(rules map[string]*Rule, part2 bool) string {
	if r == nil {
		fmt.Println("WUT? ")
		return ""
	}
	// we're a literal, just return our content
	if r.isLiteral {
		return r.contentRaw
	}
	out := strings.Builder{}
	out.WriteString("(?:")
	// We're more complex than a straight literal
	// Appears that we only ever get one OR, so split on that, then resolve down
	split := strings.Split(r.contentRaw, " | ")
	orRules := []string{}
	for _, rStr := range split {
		// resolve each rule in here, concat them together
		split := strings.Split(rStr, " ")
		concattedRules := []string{}
		for _, ruleNum := range split {
			// In part 2 we cheat, and use PCRE, slower? meh. It works.
			if part2 {
				switch ruleNum {
				case "8":
					life := rules["42"].resolve(rules, part2)
					concattedRules = append(concattedRules, fmt.Sprintf("%s+", life))
					continue
				case "11":
					life := rules["42"].resolve(rules, part2)
					notLife := rules["31"].resolve(rules, part2)
					concattedRules = append(concattedRules, fmt.Sprintf("(?<eleven>(%s%s|%[1]s(?&eleven)%[2]s))", life, notLife))
					continue

				}
			}
			toAdd := rules[ruleNum].resolve(rules, part2)
			concattedRules = append(concattedRules, toAdd)
		}
		orRules = append(orRules, strings.Join(concattedRules, ""))
	}
	out.WriteString(strings.Join(orRules, "|"))
	out.WriteRune(')')
	return out.String()
}

func rulesToReString(ruleStrs []string, allowBackreference bool, resolveNum string) string {
	rules := make(map[string]*Rule, len(ruleStrs))
	for _, v := range ruleStrs {
		split := strings.Split(v, ": ")
		rules[split[0]] = NewRule(v)
	}

	return rules[resolveNum].resolve(rules, allowBackreference)
}

func part1(input []string) string {
	reStr := rulesToReString(strings.Split(input[0], "\n"), false, "0")
	re := regexp.MustCompile("^" + reStr + "$")
	matchCount := 0
	for _, v := range strings.Split(input[1], "\n") {
		matches := re.MatchString(v)
		// fmt.Printf("%q matched by Re: %t\n", v, matches)
		if matches {
			matchCount++
		}
	}

	return fmt.Sprint(matchCount)
}

func part2(input []string) string {
	split := strings.Split(input[0], "\n")

	sort.Slice(split, func(i, j int) bool {
		return util.GetInt(strings.Split(split[i], ": ")[0]) < util.GetInt(strings.Split(split[j], ": ")[0])
	})
	split[8] = "8: 42 | 42 8"
	split[11] = "11: 42 31 | 42 11 31"

	reStr := rulesToReString(split, true, "0")
	matchCount := 0
	recurseTest := pcre.MustCompile("^"+reStr+"$", 0)

	for _, v := range strings.Split(input[1], "\n") {
		if recurseTest.MatcherString(v, 0).Matches() {
			matchCount++
		}
	}

	return fmt.Sprint(matchCount)
}

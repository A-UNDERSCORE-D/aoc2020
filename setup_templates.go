package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

const mainFile = `package main

import (
	"fmt"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.ReadLines("input.txt")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) string {
	return "stuff"
}

func part2(input []string) string {
	return "stuff2"
}
`

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var day, year int
	flag.IntVar(&day, "day", -1, "The day to create the template for")
	flag.IntVar(&year, "year", -1, "The year to create the template for")
	flag.Parse()
	if day == -1 {
		day = time.Now().Day()
		fmt.Printf("assuming day is %02d\n", day)
	}

	if year == -1 {
		year = time.Now().Year()
		fmt.Printf("assuming year is %02d\n", year)
	}

	moduleDir := fmt.Sprintf("./%d/%02d", year, day)
	panicErr(os.MkdirAll(moduleDir, 0o755))
	panicErr(os.Chdir(moduleDir))

	t := template.Must(template.New("f").Parse(mainFile))
	data := bytes.Buffer{}
	panicErr(t.Execute(&data, struct {
		Day  int
		Year int
	}{day, year}))

	panicErr(ioutil.WriteFile("solution.go", data.Bytes(), 0o600))
}

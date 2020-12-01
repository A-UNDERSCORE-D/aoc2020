package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

const mainFile = `package main

import (
	"fmt"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
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

	cookie, err := ioutil.ReadFile("cookie.txt")
	panicErr(err)

	writeTemplate(day, year)
	input := fetchInput(string(cookie), day, year)
	ioutil.WriteFile("input.txt", input, 0o600)
}

func writeTemplate(day, year int) {
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

var AoCCookieURL = func() *url.URL {
	res, err := url.Parse("https://adventofcode.com")
	panicErr(err)
	return res
}()

func fetchInput(sessionCookie string, day, year int) []byte {
	jar, err := cookiejar.New(nil)
	panicErr(err)
	jar.SetCookies(AoCCookieURL, []*http.Cookie{{Name: "session", Value: sessionCookie}})

	http.DefaultClient.Jar = jar

	res, err := http.Get(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	panicErr(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	panicErr(err)
	return data
}

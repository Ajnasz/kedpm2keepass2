package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isPwField(line string) bool {
	pwItemRegexp, _ := regexp.Compile("^(Path|Title|Username|Password|URL|Notes):")

	return pwItemRegexp.Match([]byte(line))
}

func isEmptyLine(line string) bool {
	return len(line) == 0
}

func printHeader() {
	var header []string

	header = append(header, strconv.Quote("Group"))
	header = append(header, strconv.Quote("Account"))
	header = append(header, strconv.Quote("Login Name"))
	header = append(header, strconv.Quote("Password"))
	header = append(header, strconv.Quote("Web Site"))
	header = append(header, strconv.Quote("comments"))

	fmt.Println(strings.Join(header, ","))

}

func printRows(scanner *bufio.Scanner) {
	pwItem := PwItem{}

	for scanner.Scan() {
		line := scanner.Text()

		if isEmptyLine(line) {
			if pwItem.Password != "" {
				fmt.Println(pwItem.ToCSVLine())
				pwItem = PwItem{}
			}
		} else {
			if isPwField(line) {
				pwItem.SetItemProp(line)
			}
		}
	}
}

func main() {
	var content *bufio.Scanner

	flag.Parse()
	switch name := flag.Arg(0); {
	case name == "":
		content = bufio.NewScanner(os.Stdin)

	default:
		f, err := os.Open(flag.Arg(0))

		if err != nil {
			panic(err)
		}

		defer f.Close()

		content = bufio.NewScanner(f)
	}

	printHeader()
	printRows(content)
}

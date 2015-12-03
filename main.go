package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PwItem struct {
	Title    string
	Path     string
	Username string
	Password string
	URL      string
	Notes    string
}

func fixPath(path string) string {
	if strings.Index(path, ".") == 0 {
		return path[1:]
	}

	return path
}

func (item *PwItem) ToCSVLine() string {
	var arr []string

	arr = append(arr, strconv.Quote(item.Path))
	arr = append(arr, strconv.Quote(item.Title))
	arr = append(arr, strconv.Quote(item.Username))
	arr = append(arr, strconv.Quote(item.Password))
	arr = append(arr, strconv.Quote(item.URL))
	arr = append(arr, strconv.Quote(item.Notes))

	return strings.Join(arr, ",")
}

func cleanFileContent(lines []string) []string {
	pwItemRegexp, _ := regexp.Compile("(^(Path|Title|Username|Password|URL|Notes):|^$)")

	output := make([]string, 5)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		match := pwItemRegexp.Match([]byte(line))

		if match {
			output = append(output, line)
		}
	}

	return output
}

func extractPwItems(lines []string) []PwItem {
	pwItemRegexp, _ := regexp.Compile("^(Path|Title|Username|Password|URL|Notes):")

	pwItem := PwItem{}

	output := make([]PwItem, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		match := pwItemRegexp.Match([]byte(line))

		if match {
			fieldName := line[0:strings.Index(line, ":")]
			fieldValue := strings.TrimSpace(line[strings.Index(line, ":")+1:])

			switch fieldName {
			case "Password":
				pwItem.Password = fieldValue
				break
			case "Path":
				pwItem.Path = fixPath(fieldValue)
				break
			case "Title":
				pwItem.Title = fieldValue
				break
			case "Username":
				pwItem.Username = fieldValue
				break
			case "URL":
				pwItem.URL = fieldValue
				break
			case "Notes":
				pwItem.Notes = fieldValue
				break
			}
		} else {
			if pwItem.Password != "" {
				output = append(output, pwItem)
			}
			pwItem = PwItem{}
		}
	}

	return output
}

func getContentFromStdin() []string {
	return getContent(os.Stdin)
}

func getContentFromFile(file string) []string {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	return getContent(f)

}

func getContent(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var content []string

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func main() {
	var content []string

	flag.Parse()
	switch name := flag.Arg(0); {
	case name == "":
		content = getContentFromStdin()

	default:
		content = getContentFromFile(flag.Arg(0))
	}

	cleanContent := cleanFileContent(content)

	pwItems := extractPwItems(cleanContent)

	var header []string

	header = append(header, strconv.Quote("Group"))
	header = append(header, strconv.Quote("Account"))
	header = append(header, strconv.Quote("Login Name"))
	header = append(header, strconv.Quote("Password"))
	header = append(header, strconv.Quote("Web Site"))
	header = append(header, strconv.Quote("comments"))

	fmt.Println(strings.Join(header, ","))

	for _, pwItem := range pwItems {
		fmt.Println(pwItem.ToCSVLine())
	}

}

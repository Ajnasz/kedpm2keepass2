package main

import (
	"path"
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

// password item
func (item *PwItem) GetFixedPath() string {
	newPath := path.Dir(item.Path)

	if !path.IsAbs(newPath) {
		newPath = path.Join("/", newPath)
	}

	if newPath == "/" {
		return ""
	}

	return newPath
}

func (pwItem *PwItem) SetItemProp(line string) {
	pwItemRegexp, _ := regexp.Compile("^(Path|Title|Username|Password|URL|Notes):")
	match := pwItemRegexp.Match([]byte(line))

	if match {
		fieldName := line[0:strings.Index(line, ": ")]
		fieldValue := line[strings.Index(line, ":")+2:]

		switch fieldName {
		case "Password":
			pwItem.Password = fieldValue
			break
		case "Path":
			pwItem.Path = fieldValue
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
	}
}

func (item *PwItem) ToCSVLine() string {
	var arr []string

	arr = append(arr, strconv.Quote(item.GetFixedPath()))
	arr = append(arr, strconv.Quote(item.Title))
	arr = append(arr, strconv.Quote(item.Username))
	arr = append(arr, strconv.Quote(item.Password))
	arr = append(arr, strconv.Quote(item.URL))
	arr = append(arr, strconv.Quote(item.Notes))

	return strings.Join(arr, ",")
}

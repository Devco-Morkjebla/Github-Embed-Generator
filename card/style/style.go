package style

import (
	"regexp"
)

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
	Box        string
}

func CheckHex(str map[string]string) Styles {
	var style Styles
	r, _ := regexp.Compile("^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
	if !r.MatchString(str["Border"]) {
		style.Border = "000000"
	} else {
		style.Border = str["Border"]
	}
	if !r.MatchString(str["Title"]) {
		style.Title = "000000"
	} else {
		style.Title = str["Border"]
	}
	if !r.MatchString(str["Background"]) {
		style.Background = "ffffff"
	} else {
		style.Background = str["Border"]
	}
	if !r.MatchString(str["Text"]) {
		style.Text = "000000"
	} else {
		style.Text = str["Text"]
	}
	if !r.MatchString(str["Box"]) {
		style.Box = "dddddd"
	} else {
		style.Box = str["Box"]
	}
	return style
}

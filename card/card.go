package card

import (
	"fmt"
	"githubembedapi/card/style"
	"math"
	"strings"
)

type Card struct {
	Title string       `json:"title"`
	Style style.Styles `json:"colors"`
	Body  []string     `json:"body"`
}

func (card Card) GetStyles(customStyles ...string) string {
	var style = []string{
		`<style>`,
		`.title { font: 25px sans-serif; fill: #` + card.Style.Title + `}`,
		`.text { font: 20px sans-serif; fill: #` + card.Style.Text + `; font-family: ` + card.Style.Textfont + `; text-decoration: underline;}`,
	}
	if cap(customStyles) > 0 {
		style = append(style, customStyles...)
	}

	style = append(style, `</style>`)
	return strings.Join(style, "\n")
}

func CalculateCircleProgress(progress int, radius float64) float64 {
	var c = math.Pi * (radius * 2)

	if progress < 0 {
		progress = 0
	}
	if progress > 100 {
		progress = 100
	}

	return ((100 - float64(progress)) / 100) * c
}

func GenerateCard(style style.Styles, body []string, width, height int, customStyles ...string) []string {
	var card Card
	card.Style = style
	card.GetStyles(customStyles...)
	card.Body = []string{
		fmt.Sprintf(`<svg width="%v" height="%v" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width, height, width, height),
		card.GetStyles(customStyles...),
		strings.Join(body, "\n"),
		`</svg>`,
	}
	return card.Body
}

func ToTitleCase(str string) string {
	return strings.Title(str)
}

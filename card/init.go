package card

import (
	"githubembedapi/card/style"
	"math"
	"strings"
)

type Card struct {
	Title     string       `json:"title"`
	Languages []string     `json:"languages"`
	Style     style.Styles `json:"colors"`
	Body      []string     `json:"body"`
}

func (card Card) GetStyles() {

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

func GenerateCard(style style.Styles) []string {
	var card Card
	card.Style = style
	card.GetStyles()
	test := []string{"test"}
	return test
}

func ToTitleCase(str string) string {
	return strings.Title(str)
}

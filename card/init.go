package card

import (
	"githubembedapi/card/style"
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

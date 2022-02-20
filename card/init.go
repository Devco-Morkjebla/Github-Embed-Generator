package card

import (
	"fmt"
	"strconv"
)

type CardColors struct {
	TitleColor string `json:"titleColor"`
}

type Card struct {
	Title     string   `json:"title"`
	Languages []string `json:"languages"`
	Colors    []string `json:"colors"`
	Body      []string `json:"body"`
}

func Newcard(title string, languages []string, colors []string) Card {
	totalHeight := 20
	svgTag := `<svg width="500" height="` + strconv.Itoa(totalHeight+180) + `" fill="none" viewBox="0 0 500 ` + strconv.Itoa(totalHeight+180) + `"
	xmlns="http://www.w3.org/2000/svg">`
	body := []string{
		svgTag,
		`<style>`,
		`.small { font: 20px sans-serif; fill: black}`,
		`</style>`,
		`<rect x="0" y="0" width="200" height="200" rx="15" fill="grey"/>`,
	}

	// Generate body for the languages
	for i, s := range languages {

		icon := fmt.Sprintf(`https://cdn.jsdelivr.net/gh/devicons/devicon/icons/%s/%s-original.svg`, s, s)
		img := fmt.Sprintf(`<image x="20" y="%d" href="%s" height="30" width="30"/>`, totalHeight, icon)
		text := fmt.Sprintf(`<text x="60" y="%d" class="small">%s</text>`, totalHeight+20, s)
		totalHeight += 30
		body = append(body, text)
		body = append(body, img)

		fmt.Println(i, totalHeight)
	}
	body = append(body, `</svg>`)
	newcard := Card{title, languages, colors, body}
	return newcard
}

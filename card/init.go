package card

import (
	"fmt"
	"strconv"
	"strings"
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

type Kv struct {
	Key   string
	Value int
}

type RankCard struct {
	Title  string   `json:"title"`
	Score  []Kv     `json:"score"`
	Colors []string `json:"colors"`
	Body   []string `json:"body"`
}

func Newcard(title string, languages []string, colors []string) Card {
	totalHeight := 40
	width := 400
	strokewidth := 3
	svgTag := `<svg width="` + strconv.Itoa(width+strokewidth) + `" height="` + strconv.Itoa(totalHeight+180) + `" fill="none" viewBox="0 0 ` + strconv.Itoa(width+strokewidth) + ` ` + strconv.Itoa(totalHeight+180) + `"
	xmlns="http://www.w3.org/2000/svg">`
	titlesvg := fmt.Sprintf(`<text x="20" y="25" class="large">%s</text>`, title)
	body := []string{
		svgTag,
		`<style>`,
		`.small { font: 20px sans-serif; fill: black}`,
		`.large { font: 25px sans-serif; fill: black}`,
		`</style>`,
		`<rect x="0" y="0" width="` + strconv.Itoa(width) + `" height="200" rx="15" fill="grey" style="stroke-width:3;stroke:rgba(0,0,0)"/>`,
		`<rect x="0" y="30" width="` + strconv.Itoa(width) + `" height="3" fill="black"/>`,
		titlesvg,
	}
	test := true
	if cap(languages) > 10 {
		languages = languages[:10]
	}
	// Generate body for the languages
	for i, s := range languages {
		var rowx int = 20
		var t *int = &rowx
		if totalHeight >= 180 || test == false {
			rowx = 160
			if test {
				test = false
				totalHeight = 40
			}
		}
		fmt.Println(t)
		icon := fmt.Sprintf(`https://cdn.jsdelivr.net/gh/devicons/devicon/icons/%s/%s-original.svg`, s, s)
		img := fmt.Sprintf(`<image x="%v" y="%v" href="%v" height="30" width="30"/>`, rowx, totalHeight, icon)
		text := fmt.Sprintf(`<text x="%v" y="%v" class="small">%v</text>`, rowx+40, totalHeight+20, s)
		totalHeight += 30

		// panic("This crappy code causes error")
		body = append(body, text)
		body = append(body, img)

		fmt.Println(i, totalHeight)
	}
	body = append(body, `</svg>`)
	newcard := Card{title, languages, colors, body}
	return newcard
}

func Rankcard(title string, score []Kv, colors []string) RankCard {
	totalHeight := 40
	width := 400
	strokewidth := 3
	svgTag := `<svg width="` + strconv.Itoa(width+strokewidth) + `" height="` + strconv.Itoa(totalHeight+180) + `" fill="none" viewBox="0 0 ` + strconv.Itoa(width+strokewidth) + ` ` + strconv.Itoa(totalHeight+180) + `"
	xmlns="http://www.w3.org/2000/svg">`
	titlesvg := fmt.Sprintf(`<text x="20" y="25" class="large">%s</text>`, title)
	body := []string{
		svgTag,
		`<style>`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.small { font: 20px sans-serif; fill: black; font-family: Papyrus;}`,
		`.large { font: 25px sans-serif; fill: black}`,
		`</style>`,
		`<rect x="0" y="0" width="` + strconv.Itoa(width) + `" height="200" rx="15" fill="white" style="stroke-width:3;stroke:rgba(0,0,0)"/>`,
		`<rect x="0" y="30" width="` + strconv.Itoa(width) + `" height="3" fill="black"/>`,
		titlesvg,
	}
	// Generate body for the languages
	for _, kv := range score {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
	pos := 1
	for _, s := range score {
		var rowx int = 20

		text := fmt.Sprintf(`<text x="%v" y="%v" class="small">%v. %v - %v commits</text>`, rowx, totalHeight+20, pos, s.Key, s.Value)
		totalHeight += 30
		pos += 1
		body = append(body, text)

	}
	body = append(body, `</svg>`)
	newcard := RankCard{ToTitleCase(title), score, colors, body}
	return newcard
}
func ToTitleCase(str string) string {
	return strings.Title(str)
}

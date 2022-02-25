package card

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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
type User struct {
	Avatar string
	Score  int
	Name   string
}
type Kv struct {
	Key   string
	Value User
}

type RankCard struct {
	Title  string   `json:"title"`
	Score  []Kv     `json:"score"`
	Styles Styles   `json:"styles"`
	Body   []string `json:"body"`
}
type Response struct {
	Total_Count int     `json:"total_count"`
	Items       []Items `json:"items"`
}
type Items struct {
	Url          string `json:"url"`
	Comments_url string `json:"comments_url"`
	Author       Author `json:"author"`
}
type Author struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
}

func Newcard(title string, languages []string, colors []string) Card {
	totalHeight := 40
	width := 400
	strokewidth := 3
	svgTag := `<svg width="` + strconv.Itoa(width+strokewidth) + `" height="` + strconv.Itoa(totalHeight+180) + `" fill="none" viewBox="0 0 ` + strconv.Itoa(width+strokewidth) + ` ` + strconv.Itoa(totalHeight+180) + `"
	xmlns="http://www.w3.org/2000/svg">`
	titlesvg := fmt.Sprintf(`<text x="20" y="25" class="title">%s</text>`, title)
	body := []string{
		svgTag,
		`<style>`,
		`.small { font: 20px sans-serif; fill: black}`,
		`.large { font: 25px sans-serif; fill: black}`,
		`.title { font: 25px sans-serif; fill: black}`,
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

func Rankcard(title string, users []string, style Styles) RankCard {

	ss := make(map[string]User)
	for key, i := range users {
		userurl := "https://api.github.com/search/commits?q=author:" + fmt.Sprintf("%v", i) + "&sort=author-date&order=desc&page=1"
		response, err := http.Get(userurl)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		var responseObject Response
		decodeerr := json.Unmarshal(responseData, &responseObject)

		if decodeerr != nil {
			panic(decodeerr)
		}

		ss[fmt.Sprintf("%v", users[key])] = User{Avatar: responseObject.Items[0].Author.Avatar_Url, Score: responseObject.Total_Count, Name: responseObject.Items[0].Author.Login}
	}

	// Sort Scores
	var score []Kv
	for k, v := range ss {
		score = append(score, Kv{k, v})
	}

	// Sort score
	sort.Slice(score, func(i, j int) bool {
		return score[i].Value.Score > score[j].Value.Score
	})

	totalHeight := 40
	width := 400
	strokewidth := 3
	svgTag := `<svg width="` + strconv.Itoa(width+strokewidth) + `" height="` + strconv.Itoa(totalHeight+180) + `" fill="none" viewBox="0 0 ` + strconv.Itoa(width+strokewidth) + ` ` + strconv.Itoa(totalHeight+180) + `"
	xmlns="http://www.w3.org/2000/svg">`

	titlesvg := fmt.Sprintf(`<text x="20" y="25" class="title">%s</text>`, ToTitleCase(title))
	body := []string{
		svgTag,
		`<style>`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: ` + style.Text + `; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.large { font: 25px sans-serif; fill: black}`,
		`.title { font: 25px sans-serif; fill: ` + style.Title + `}`,
		`.box { fill: ` + style.Background + `}`,
		`.profileimage { border-radius: 50%}`,
		`</style>`,
		`<rect x="0" y="0" class="box" width="` + strconv.Itoa(width) + `" height="200" rx="15" style="stroke-width:3;stroke:` + style.Border + `"/>`,
		`<rect x="0" y="30" width="` + strconv.Itoa(width) + `" height="3" fill="` + style.Border + `"/>`,
		titlesvg,
	}
	// Generate body for the users
	pos := 1
	for _, s := range score {
		var rowx int = 20

		img := fmt.Sprintf(`<image x="%v" y="%v" href="%v" class="profileimage" height="30" width="30"/>`, rowx, totalHeight, s.Value.Avatar)
		text := fmt.Sprintf(`<text x="%v" y="%v" class="text">%v. %v - %v commits</text>`, rowx+40, totalHeight+20, pos, ToTitleCase(s.Value.Name), s.Value.Score)
		totalHeight += 30
		pos += 1
		body = append(body, text)
		body = append(body, img)

	}
	body = append(body, `</svg>`)
	newcard := RankCard{title, score, style, body}
	return newcard
}
func ToTitleCase(str string) string {
	return strings.Title(str)
}

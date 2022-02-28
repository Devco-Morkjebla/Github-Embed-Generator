package skills

import (
	"fmt"
	"strconv"
	"strings"
)

type Skillscard struct {
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
	Styles Styles   `json:"styles"`
	Body   []string `json:"body"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
	Box        string
}

func Skills(title string, languages []string, style Styles) Skillscard {

	height := 700
	width := 600
	titleboxheight := 50
	padding := 10
	strokewidth := 3
	boxwidth := 100
	boxheight := 100
	body := []string{
		`<script type="application/ecmascript"> <![CDATA[
			function skill_hover(evt) {
			  var txt = document.getElementById("lang");
			  var skill = evt.target;
			  var currentSkill = skill.getAttribute("title");
			  txt.textContent = ToTitleCase(currentSkill);
			}
			function ToTitleCase(string) {
				return string.charAt(0).toUpperCase() + string.slice(1)
			}

		]]> </script>`,
		`<style>`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: #` + style.Text + `; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.large {
			font: 25px sans-serif; 
			fill: black
		}`,
		`.title { font: 25px sans-serif; fill: #` + style.Title + `}`,
		`.repobox { 
			fill: #` + style.Box + `;
			border: ` + strconv.Itoa(strokewidth) + `px solid #` + style.Border + `;
		}`,
		`.repobox:hover { fill: rgba(255,0,0,0.8);}`,
		`.repobox:hover rect {stroke-width: ` + strconv.Itoa(strokewidth+3) + `px;}`,
		`.box {
			fill: #` + style.Background + `;
			border: 3px solid #` + style.Border + `;
			stroke: #` + style.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
		`</style>`,
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, ToTitleCase(title)),
	}
	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding

	posX := 0
	imgsize := 80
	originalpos := posX
	newwidth := width
	row := func(content []string, lang string) {

		bodyAdd(fmt.Sprintf(`<g class="repobox" onmouseenter="skill_hover(evt)" title="%v" transform="translate(%v,%v) rotate(0)">`, lang, posX+padding, posY))
		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		// check if next box will fit into card
		if posX+boxwidth+(boxwidth+padding) >= width {
			posY += boxheight + padding
			newwidth = posX + boxwidth + (padding * 2)
			posX = originalpos - (boxwidth + padding)
		}
	}
	for _, v := range languages {

		icon := fmt.Sprintf(`https://cdn.jsdelivr.net/gh/devicons/devicon/icons/%v/%v-original.svg`, v, v)
		if v == "tailwindcss" {
			icon = fmt.Sprintf(`https://cdn.jsdelivr.net/gh/devicons/devicon/icons/%v/%v-plain.svg`, v, v)
		}
		img := fmt.Sprintf(`<image x="%v" y="%v" href="%v" height="%v" width="%v"/>`, boxwidth-imgsize-padding, boxheight-imgsize-padding, icon, imgsize, imgsize)

		row([]string{
			fmt.Sprintf(`<rect x="0" y="0" rx="5" class="box" width="%v" height="%v" />`, boxwidth, boxheight),
			img,
		}, v)

		posX += boxwidth + padding
	}
	if newwidth != width {
		width = newwidth
	}
	// Bottom / Footer
	langstring := "None"

	bodyAdd(fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="#%v"/>`, height-titleboxheight, width, strokewidth, style.Border))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" id="lang" class="text">%v</text>`, (width/2)-((len(langstring)*20)/2), (height - (titleboxheight / 2)), langstring))

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="#%v"/>`, titleboxheight, width, strokewidth, style.Border)}, body...)
	body = append([]string{fmt.Sprintf(`<rect x="0" y="0" class="box" width="%v" height="%v" rx="15"  />`, width, height)}, body...)
	svgTag := fmt.Sprintf(`<svg width="%v" height="%v" fill="none" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width+strokewidth, height+strokewidth, width+strokewidth, height+strokewidth)
	body = append([]string{svgTag}, body...)
	bodyAdd(`</svg>`)
	newcard := Skillscard{title, languages, style, body}
	return newcard

}
func ToTitleCase(str string) string {
	return strings.Title(str)
}

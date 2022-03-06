package skills

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/icons"
	"math"
	"strconv"
	"strings"
)

type Skillscard struct {
	Title  string       `json:"title"`
	Skills []string     `json:"skills"`
	Styles style.Styles `json:"styles"`
	Body   []string     `json:"body"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
	Box        string
}

func Skills(title string, languages []string, style style.Styles) Skillscard {

	height := 700
	width := 600
	titleboxheight := 50
	padding := 10
	strokewidth := 3
	boxwidth := 60
	boxheight := 60
	body := []string{
		`<style>`,
		`@keyframes gradient {
			0% {
				background-position: 0% 50%;
			}
			50% {
				background-position: 100% 50%;
			}
			100% {
				background-position: 0% 50%;
			}
		}`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: #` + style.Text + `; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.textwhite { font: 20px sans-serif; fill: #ffffff; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.large {
			font: 25px sans-serif; 
			fill: black
		}`,
		`.title { 
			font: 25px sans-serif; 
			fill: #` + style.Title + `;
		}`,
		`.repobox { 
			fill: #` + style.Box + `;
			border: ` + strconv.Itoa(strokewidth) + `px solid #` + style.Border + `;
		}`,
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

	// Algoritm for checking if color is too dark
	colorToDark := func(color string) bool {
		var c = strings.Replace(color, "#", "", -1) // strip #
		rgb, err := strconv.ParseInt(c, 16, 32)     // convert rrggbb to decimal
		if err != nil {
			panic(err.Error())
		}
		r := (rgb >> 16) & 0xff // extract red
		g := (rgb >> 8) & 0xff  // extract green
		b := (rgb >> 0) & 0xff  // extract blue

		rFloat := 0.2126
		gFloat := 0.7152
		bFloat := 0.0722
		r2Float := float64(r)
		g2Float := float64(g)
		b2Float := float64(b)
		luma := math.Sqrt(rFloat*(r2Float*r2Float) +
			gFloat*(g2Float*g2Float) +
			bFloat*(b2Float*b2Float))

		return luma < 80
	}
	fmt.Println(colorToDark("#000000"))
	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding

	posX := 0

	imgsize := boxwidth - (padding * 2)
	originalpos := posX
	newwidth := width
	newheight := height

	row := func(content []string, lang string) {
		bodyAdd(fmt.Sprintf(`<g class="repobox" title="%v" transform="translate(%v,%v) rotate(0)">`, lang, posX+padding, posY))

		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		newheight = posY + boxheight + padding
		// check if next box will fit into card
		if posX+boxwidth+(boxwidth+padding) >= width {
			posY += boxheight + padding
			newwidth = posX + boxwidth + (padding * 2)
			posX = originalpos - (boxwidth + padding)
		}
	}

	for _, v := range languages {

		icon := icons.Icons(v)
		img := fmt.Sprintf(`<g x="%v" y="%v" height="%v" width="%v">%v</g>`, boxwidth-imgsize, boxheight-imgsize, imgsize, imgsize, icon)

		row([]string{
			fmt.Sprintf(`<rect x="0" y="0" rx="5" class="" width="%v" height="%v" />`, boxwidth, boxheight),
			img,
		}, v)

		posX += boxwidth + padding
	}

	// adjust the svg size to the content
	if newwidth != width {
		width = newwidth
	}
	if newheight != height {
		height = newheight
	}

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="#%v"/>`, titleboxheight, width, strokewidth, style.Border)}, body...)
	body = append([]string{fmt.Sprintf(`<rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height)}, body...)
	svgTag := fmt.Sprintf(`<svg width="%v" height="%v" fill="none" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width+strokewidth, height+strokewidth, width+strokewidth, height+strokewidth)
	body = append([]string{svgTag}, body...)
	bodyAdd(`</svg>`)
	newcard := Skillscard{title, languages, style, body}
	return newcard

}
func ToTitleCase(str string) string {
	return strings.Title(str)
}

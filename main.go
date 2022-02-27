package main

import (
	"fmt"
	"githubembedapi/card"
	"githubembedapi/organization"
	"net/http"
	"os"
	"strings"

	"regexp"

	"github.com/alexsasharegan/dotenv"
	"github.com/gin-gonic/gin"
)

//https://go.dev/doc/tutorial/web-service-gin

func main() {
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/ranklist", rankList)
	router.GET("/skills", getSkills)
	router.GET("/card", getCard)
	router.GET("/mostactivity", getMostactivity)
	// router.Run("localhost:8080")
	router.Run()

}

func getCard(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	colors := []string{"red", "blue"}
	languages := strings.Split(c.Request.URL.Query().Get("languages"), ",")
	title := c.Request.URL.Query().Get("title")
	if len(title) > 0 {
		title = "Languages"
	}
	newCard := card.Newcard(title, languages, colors)

	fmt.Println(newCard)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}
func getMostactivity(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color organization.Styles
	org := c.Request.FormValue("org")
	title := c.Request.FormValue("title")
	bordercolor := c.Request.FormValue("bordercolor")
	titlecolor := c.Request.FormValue("titlecolor")
	backgroundcolor := c.Request.FormValue("backgroundcolor")
	textcolor := c.Request.FormValue("textcolor")
	textfont := c.Request.FormValue("textfont")
	boxcolor := c.Request.FormValue("boxcolor")
	if title == "" {
		title = "Rank"
	}

	r, _ := regexp.Compile("^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
	if !r.MatchString(bordercolor) {
		bordercolor = "000000"
	}
	if !r.MatchString(titlecolor) {
		titlecolor = "000000"
	}
	fmt.Println(r.MatchString(backgroundcolor))
	if !r.MatchString(backgroundcolor) {
		backgroundcolor = "ffffff"
	}
	if !r.MatchString(textcolor) {
		textcolor = "000000"
	}
	if textfont == "" {
		textfont = "Helvetica"
	}
	if !r.MatchString(boxcolor) {
		boxcolor = "000000"
	}
	color.Border = bordercolor
	color.Title = titlecolor
	color.Background = backgroundcolor
	color.Text = textcolor
	color.Textfont = textfont
	color.Box = boxcolor

	github_token := os.Getenv("GITHUB")
	newCard := organization.MostactivityCard(title, org, color, github_token)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}
func rankList(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color card.Styles
	users := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("users")), ",")
	title := c.Request.FormValue("title")
	bordercolor := c.Request.FormValue("bordercolor")
	titlecolor := c.Request.FormValue("titlecolor")
	backgroundcolor := c.Request.FormValue("backgroundcolor")
	textcolor := c.Request.FormValue("textcolor")
	textfont := c.Request.FormValue("textfont")

	if title == "" {
		title = "Rank"
	}
	if len(users) > 5 {
		users = users[:5]
	}
	if bordercolor == "" {
		bordercolor = "black"
	}
	if titlecolor == "" {
		titlecolor = "black"
	}
	if backgroundcolor == "" {
		backgroundcolor = "white"
	}
	if textcolor == "" {
		textcolor = "black"
	}
	if textfont == "" {
		textfont = "Helvetica"
	}
	color.Border = bordercolor
	color.Title = titlecolor
	color.Background = backgroundcolor
	color.Text = textcolor
	color.Textfont = textfont
	newCard := card.Rankcard(title, users, color)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}

func getSkills(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	languages := strings.Split(c.Request.URL.Query().Get("languages"), ",")

	body := []string{`<svg width="500" height="500" fill="none" viewBox="0 0 500 500"
	xmlns="http://www.w3.org/2000/svg">`}

	// Generate body for the languages
	for i, s := range languages {
		text := fmt.Sprintf(`<text x="30" y="%d">%s</text>`, i*10, s)
		body = append(body, text)
	}
	body = append(body, `</svg>`)

	fmt.Println(strings.Join(body, "\n"))
	var svg string = `<svg width="500" height="500" fill="none" viewBox="0 0 500 500"
	xmlns="http://www.w3.org/2000/svg">
	<style>
        .small { font: 20px sans-serif; fill: black}
        .heavy { font: bold 30px sans-serif; }

    </style>
	<defs>
		<linearGradient gradientTransform="rotate(30)" id="redbluegreenpurple">
			<stop stop-color="red" offset="0%"/>
			<stop stop-color="blue" offset="25%"/>
			<stop stop-color="green" offset="50%"/>
			<stop stop-color="purple" offset="75%"/>
		</linearGradient>
	</defs>
	<rect x="0" y="0" width="200" height="200" rx="15" fill="url(#redbluegreenpurple)"/>
	<image x="20" y="20" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/php/php-original.svg" height="30" width="30"/>
	<text x="60" y="40" class="small">php</text>
	<image x="20" y="50" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="30" width="30"/>
	<text x="60" y="70" class="small">go</text>
	<image x="20" y="80" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/typescript/typescript-original.svg" height="30" width="30"/>
	<text x="60" y="100" class="small">typescript</text>
	<image x="20" y="110" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mysql/mysql-original.svg" height="30" width="30"/>
	<text x="60" y="130" class="small">mysql</text>
	<image x="20" y="140" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/javascript/javascript-original.svg" height="30" width="30"/>
	<text x="60" y="160" class="small">javascript</text>
</svg>
    `

	c.String(http.StatusOK, svg)
}

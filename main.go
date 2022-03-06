package main

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/organization"
	"githubembedapi/project"
	"githubembedapi/rank"
	"net/http"
	"strings"

	"regexp"

	"githubembedapi/skills"

	"github.com/gin-gonic/gin"
)

type Icons []struct {
	Name string `json:"name"`
	Svg  string `json:"svg"`
}

func main() {

	router := gin.Default()
	router.GET("/ranklist", rankList)
	router.GET("/skills", getSkills)
	router.GET("/mostactivity", getMostactivity)
	router.GET("/project", projectcard)
	router.Run("localhost:8080")
	// router.Run()

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

	// github_token := os.Getenv("GITHUB")
	github_token := ""
	newCard := organization.MostactivityCard(title, org, color, github_token)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}
func projectcard(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}
	repo := c.Request.FormValue("repo")
	color = style.CheckHex(styles)
	c.String(http.StatusOK, project.Project(repo, color))
}
func rankList(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color rank.Styles
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
	newCard := rank.Rankcard(title, users, color)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}

func getSkills(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	// Define styles
	var color style.Styles
	languages := strings.Split(c.Request.URL.Query().Get("languages"), ",")

	styles := map[string]string{
		color.Title:      c.Request.FormValue("titlecolor"),
		color.Border:     c.Request.FormValue("bordercolor"),
		color.Background: c.Request.FormValue("backgroundcolor"),
		color.Text:       c.Request.FormValue("textcolor"),
		color.Box:        c.Request.FormValue("boxcolor"),
	}

	color = style.CheckHex(styles)
	title := c.Request.FormValue("title")

	if title == "" {
		title = "Skills"
	}

	newCard := skills.Skills(title, languages, color)

	c.String(http.StatusOK, strings.Join(newCard.Body, "\n"))
}

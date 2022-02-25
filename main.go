package main

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//https://go.dev/doc/tutorial/web-service-gin

func main() {
	router := gin.Default()
	router.GET("/skills", getSkills)
	router.GET("/card", getCard)
	router.GET("/rankList", rankList)
	router.Run("localhost:8080")

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }
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

type Response struct {
	Total_Count int     `json:"total_count"`
	Items       []Items `json:"items"`
}
type Items struct {
	Url          string `json:"url"`
	Comments_url string `json:"comments_url"`
}

func rankList(c *gin.Context) {
	// c.Header("Content-Type", "image/svg+xml")

	users := strings.Split(c.Request.URL.Query().Get("brukere"), ",")

	fmt.Println(users)
	for key := range users {
		fmt.Println(key)
		userurl := "https://api.github.com/search/commits?q=author:" + fmt.Sprintf("%v", key) + "&sort=author-date&order=desc&page=1"
		// userurl := "https://api.github.com/search/commits?q=author:lartrax&sort=author-date&order=desc&page=1"

		fmt.Println(key)
		response, err := http.Get(userurl)

		// fmt.Println(response.Body)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(response.Body)

		var responseObject Response
		// decodeerr := json.Unmarshal(responseData, &responseObject)
		decodeerr := decoder.Decode(&responseObject)

		if decodeerr != nil {
			panic(decodeerr)
		}
		fmt.Println(responseObject)
		c.String(http.StatusOK, string(responseObject.Total_Count))
	}

	// title := "test"
	// newCard := card.Newcard(title, users, colors)

	// fmt.Println(newCard)

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

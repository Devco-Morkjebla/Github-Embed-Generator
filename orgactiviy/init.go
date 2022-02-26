package organization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Repos []Repo `json:""`
}
type Repo struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string
}
type Author struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
}

type OrgCard struct {
	Title        string   `json:"title"`
	Organization string   `json:"score"`
	Styles       Styles   `json:"styles"`
	Body         []string `json:"body"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
}

func MostactivityCard(title string, org string, style Styles) OrgCard {
	userurl := "https://api.github.com/orgs/" + org + "/repos"
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

}

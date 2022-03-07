package commit_activity

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RepoActivity []struct {
	Total int   `json:"total"`
	Week  int   `json:"week"`
	Days  []int `json:"days"`
}

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func RepositoryCommitActivity(title, user, repo string, cardstyle style.Styles) string {
	apiurl := "https://api.github.com/repos/" + user + "/" + repo + "/stats/commit_activity"

	reqAPI, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		panic(err.Error())
	}
	clientAPI := &http.Client{}

	responseAPI, err := clientAPI.Do(reqAPI)
	defer recoverFromError()
	if err != nil {
		panic(err.Error())
	}
	defer responseAPI.Body.Close()

	responseDataAPI, err := ioutil.ReadAll(responseAPI.Body)
	if err != nil {
		panic(err)
	}

	var resObjectAPI RepoActivity
	json.Unmarshal(responseDataAPI, &resObjectAPI)

	start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC)

	fmt.Println(start)
	fmt.Println(end)

	// calculate total number of days
	duration := end.Sub(start)
	daysOfYear := int(duration.Hours() / 24)

	fmt.Printf("difference %d days", daysOfYear)

	customstyles := []string{
		`.circle {
		transform: rotate(-90deg);
		}`,
		`.rank-circle-rim {
			stroke: #333333;
			fill: none;
			opacity: 0.4;	
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}
	paddingX := 30
	paddingY := 30

	body := []string{
		`<g id="Box">`,
		fmt.Sprintf(`<rect x="0" y="0" rx="15" width="%v" height="%v" />`, 800, 300),
		`</g>`,
		`<g data-testid="card-text">`,
		fmt.Sprintf(`<text x="%v" y="%v" id="Stats" class="title">%v Stats</text>`, paddingX, paddingY, card.ToTitleCase(repo)),
		fmt.Sprintf(`<line id="gradLine" x1="%v" y1="40" x2="400" y2="40" stroke="url(#paint0_angular_0_1)"/>`, paddingX),
		`</g>`,
	}
	gridX := 30
	gridY := 50
	gridPadding := 2
	gridBoxSize := 10

	grid := []string{`<g data-testid="card-grid">`}
	for week, data := range resObjectAPI {
		grid = append(grid, fmt.Sprintf(`<g id="week%v">`, week))
		for days, commits := range data.Days {
			color := "#002400"
			if commits <= 10 {
				color = "#002400"
			} else if commits <= 20 {
				color = "#005700"
			} else if commits <= 30 {
				color = "#008a00"
			} else if commits <= 40 {
				color = "#52b152"
			} else if commits <= 50 {
				color = "#83c783"
			} else if commits >= 60 {
				color = "#b4ddb4"
			}

			grid = append(grid, fmt.Sprintf(`<rect id="day%v" fill="%v" width="%v" height="%v" x="%v" y="%v"></rect>`, days, color, gridBoxSize, gridBoxSize, gridX, gridY))
			// grid = append(grid, fmt.Sprintf(`<text id="day%v" class="text" x="%v" y="%v" fill="red">%v</text>`, days, gridX, gridY, commits))
			gridY += gridBoxSize + gridPadding
		}
		grid = append(grid, `</g>`)
		gridX += gridBoxSize + gridPadding
		gridY = 50
	}
	grid = append(grid, `</g>`)
	body = append(body, strings.Join(grid, "\n"))
	return strings.Join(card.GenerateCard(cardstyle, defs, body, 800, 300, customstyles...), "\n")
}

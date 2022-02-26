package organization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxID string `json:"spdx_id"`
		URL    string `json:"url"`
		NodeID string `json:"node_id"`
	} `json:"license"`
	AllowForking  bool          `json:"allow_forking"`
	IsTemplate    bool          `json:"is_template"`
	Topics        []interface{} `json:"topics"`
	Visibility    string        `json:"visibility"`
	Forks         int           `json:"forks"`
	OpenIssues    int           `json:"open_issues"`
	Watchers      int           `json:"watchers"`
	DefaultBranch string        `json:"default_branch"`
	Permissions   struct {
		Admin    bool `json:"admin"`
		Maintain bool `json:"maintain"`
		Push     bool `json:"push"`
		Triage   bool `json:"triage"`
		Pull     bool `json:"pull"`
	} `json:"permissions"`
}
type Content []struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Sha         string      `json:"sha"`
	Size        int         `json:"size"`
	URL         string      `json:"url"`
	HTMLURL     string      `json:"html_url"`
	GitURL      string      `json:"git_url"`
	DownloadURL interface{} `json:"download_url"`
	Type        string      `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
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
	Box        string
}

func MostactivityCard(title string, org string, style Styles) OrgCard {

	// userurl := "https://api.github.com/orgs/" + org + "/repos"
	userurl := "https://api.github.com/orgs/devco-morkjebla/repos"
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
	json.Unmarshal(responseData, &responseObject)

	height := 600
	width := 600
	padding := 10
	strokewidth := 3
	svgTag := `<svg width="` + strconv.Itoa(width+strokewidth) + `" height="` + strconv.Itoa(height+strokewidth) + `" fill="none" viewBox="0 0 ` + strconv.Itoa(width+strokewidth) + ` ` + strconv.Itoa(height+strokewidth) + `"
	xmlns="http://www.w3.org/2000/svg">`
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
		`<rect x="0" y="0" class="box" width="` + strconv.Itoa(width) + `" height="` + strconv.Itoa(height) + `" rx="15" style="stroke-width:3;stroke:` + style.Border + `"/>`,
		`<rect x="0" y="30" width="` + strconv.Itoa(width) + `" height="3" fill="` + style.Border + `"/>`,
		fmt.Sprintf(`<text x="20" y="25" class="title">%s</text>`, ToTitleCase(org)),
	}
	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}
	row := func(content []string, posX int, posY int) string {
		for _, v := range content {

			fields := strings.Fields(v)

			// Set X pos
			fields[1] = strings.Join([]string{`x="`, strconv.Itoa(posX), `"`}, "")
			// Set Y pos
			fields[2] = strings.Join([]string{`y="`, strconv.Itoa(posY), `"`}, "")

			bodyAdd(strings.Join(fields, ""))
			posY += 30
		}
		return "test"
	}
	// fmt.Println(row("s"))
	fmt.Println(len(responseObject))
	for _, r := range responseObject {
		response2, err := http.Get("https://api.github.com/repos/" + org + "/CameraLight/contents/")
		if err != nil {
			panic(err.Error())
		}
		responseData2, err := ioutil.ReadAll(response2.Body)
		if err != nil {
			panic(err)
		}
		var responseContent Content
		json.Unmarshal(responseData2, &responseContent)
		pos := 80
		row([]string{
			`<rect x="0" y="" width="270" height="100" fill="black"/>`,
			`<text x="" y="" class="text">` + r.Name + `</text>`,
			`<text x="" y="" class="text">Language - ` + r.Language + `</text>`,
			`<text x="" y="" class="text">Issues - ` + strconv.Itoa(r.OpenIssuesCount) + `</text>`,
			`<text x="" y="" class="text">Files - ` + strconv.Itoa(len(responseContent)) + `</text>`,
		}, padding*2, pos)
		pos += 100
	}
	bodyAdd(`</svg>`)

	newcard := OrgCard{title, org, style, body}
	return newcard

}
func ToTitleCase(str string) string {
	return strings.Title(str)
}

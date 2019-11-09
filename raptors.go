package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

// Team playing
type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	Nickname     string `json:"nickname"`
	URLName      string `json:"url_name"`
	Score        string `json:"score"`
}

// Period, time, game status
type PeriodTime struct {
	GameStatus  string `json:"game_status"`
	Period      string `json:"period_value"`
	Clock       string `json:"game_clock"`
}

// GameRecord game scheduled
type GameRecord struct {
	ID               int        `json:"id"`
	Date             string     `json:"date"`
	Time             string     `json:"time"`
	Arena            string     `json:"arena"`
	City             string     `json:"city"`
	State            string     `json:"state"`
	Country          string     `json:"country"`
	HomeStartDate    string     `json:"home_start_date"`
	HomeStartTime    string     `json:"home_start_time"`
	VisitorStartDate string     `json:"visitor_start_date"`
	VisitorStartTime string     `json:"visitor_start_time"`
	Home             Team       `json:"home"`
	Visitor          Team       `json:"visitor"`
	Period           int        `json:"period"`
	Postseason       bool       `json:"postseason"`
	Season           int        `json:"season"`
	Status           PeriodTime `json:"period_time"`
	VisitorTeamScore int        `json:"visitor_team_score"`
	IsHomeTeam       int        `json:"is_home_team"`
	Outcome          string     `json:"outcome"`
}

// Response - JSON Body
type Response struct {
	Game []GameRecord `json:"game"`
}

func processResponse(body []byte) [][]interface{} {
	var result map[string]Response
	json.Unmarshal([]byte(body), &result)
	var output [][]interface{}
	for _, game := range result["sports_content"].Game {
		startDate, _ := time.ParseInLocation("20060102", game.HomeStartDate, time.Local)
		var startTime string
		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
		if now.Before(startDate) || now.Equal(startDate) {
			playingWith := ""
			hommies := ""
			starts := ""
			status := game.Status.GameStatus
			if game.IsHomeTeam == 1 {
				playingWith = game.Visitor.Nickname
				hommies = game.Home.Nickname
				startTime = game.HomeStartTime
			} else {
				playingWith = game.Home.Nickname
				hommies = game.Visitor.Nickname
				startTime = game.VisitorStartTime
			}
			if status == "1" {
				starts = fmt.Sprintf("%s %s:%s",
					startDate.Format("Nov 02, 2006"), startTime[0:2], startTime[2:4])
			} else if status == "2" {
				starts = fmt.Sprintf("Q%s - %s", game.Status.Period, game.Status.Clock)
			} else {
				starts = "Final"
			}
			teams := fmt.Sprintf("%s vs %s", hommies, playingWith)
			place := fmt.Sprintf("%s, %s, %s, %s", game.Arena, game.City, game.State, game.Country)
			output = append(output, []interface{}{starts, teams, place})
		}
		if len(output) == 5 {
			break
		}
	}
	return output
}

func main() {
	teamSlug := "raptors"
	year := 2019
	url := fmt.Sprintf("http://data.nba.net/json/cms/%d/team/%s/schedule.json", year, teamSlug)

	req, _ := http.NewRequest("GET", url, nil)
	res, error := http.DefaultClient.Do(req)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while connecting to API. Please try again later: %v\n", error)
		os.Exit(1)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	output := processResponse(body)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"DATE", "TEAMS", "PLACE"})
	for _, line := range output {
		t.AppendRow(line)
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

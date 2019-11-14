package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bilogub/raptors/pkg/renderer"
	"github.com/bilogub/raptors/pkg/response"
)

func main() {
	showPreviousGames := flag.Bool("prev", false, "Show previous 5 games")
	teamSlug := flag.String("team", "raptors", "Pick your team. Default is raptors")
	recordsNumToShow := flag.Int("num", 5, "How many records to show. Default is 5")
	flag.Parse()
	year := time.Now().Year()
	url := fmt.Sprintf("http://data.nba.net/json/cms/%d/team/%s/schedule.json", year, *teamSlug)

	req, _ := http.NewRequest("GET", url, nil)
	res, error := http.DefaultClient.Do(req)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while connecting to API. Please try again later: %v\n", error)
		os.Exit(1)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	output := response.Processor{}.Call(body, *showPreviousGames, *recordsNumToShow)

	renderer.Renderer{}.Call(renderer.Terminal{}, output)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bilogub/raptors/pkg/renderer"
	"github.com/bilogub/raptors/pkg/response"
)

func main() {
	teamSlug := "raptors"
	year := time.Now().Year()
	url := fmt.Sprintf("http://data.nba.net/json/cms/%d/team/%s/schedule.json", year, teamSlug)

	req, _ := http.NewRequest("GET", url, nil)
	res, error := http.DefaultClient.Do(req)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while connecting to API. Please try again later: %v\n", error)
		os.Exit(1)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	output := response.Processor{}.Call(body)

	renderer.Renderer{}.Call(renderer.Terminal{}, output)
}

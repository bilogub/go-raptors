package response

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Processor parses the JSON response and returns structs
type Processor struct{}

func (p Processor) homeTeam(game GameRecord) string {
	if game.IsHomeTeam == "1" {
		return game.Home.Nickname
	}
	return game.Visitor.Nickname
}

func (p Processor) guestTeam(game GameRecord) string {
	if game.IsHomeTeam == "1" {
		return game.Visitor.Nickname
	}
	return game.Home.Nickname
}

func (p Processor) gameStarts(game GameRecord) string {
	if game.IsHomeTeam == "1" {
		return game.HomeStartTime
	}
	return game.VisitorStartTime
}

func (p Processor) formattedStartTime(game GameRecord) string {
	switch game.Status.GameStatus {
	case "1", "3":
		startTime := p.gameStarts(game)
		startDate, _ := time.ParseInLocation("20060102", game.HomeStartDate, time.Local)
		return fmt.Sprintf("%s %s:%s", startDate.Format("Jan 02, 2006"), startTime[0:2], startTime[2:4])
	case "2":
		return fmt.Sprintf("Q%s - %s", game.Status.Period, game.Status.Clock)
	default:
		return ""
	}
}

func (p Processor) formattedScore(game GameRecord) string {
	status := game.Status.GameStatus
	if status != "2" && status != "3" {
		return ""
	}

	if game.IsHomeTeam == "1" {
		return fmt.Sprintf("%s - %s", game.Home.Score, game.Visitor.Score)
	}
	return fmt.Sprintf("%s - %s", game.Visitor.Score, game.Home.Score)
}

func (p Processor) formattedTeams(game GameRecord) string {
	return fmt.Sprintf("%s vs %s", p.homeTeam(game), p.guestTeam(game))
}

func (p Processor) formattedPlace(game GameRecord) string {
	return fmt.Sprintf("%s, %s, %s, %s", game.Arena, game.City, game.State, game.Country)
}

// Call returns strctured data from API request response
func (p Processor) Call(body []byte, showPrevious bool, recordsNumToShow int) [][]interface{} {
	var result map[string]Response
	error := json.Unmarshal([]byte(body), &result)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while trying to parse the response. Please try again later: %v\n", error)
		os.Exit(1)
	}

	var output [][]interface{}
	for i, game := range result["sports_content"].Game {
		startDate, _ := time.ParseInLocation("20060102", game.HomeStartDate, time.Local)

		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)

		if now.Before(startDate) || now.Equal(startDate) {
			previousGames := []GameRecord{}
			if showPrevious {
				backward := i - recordsNumToShow
				if backward < 0 {
					backward = 0
				}
				previousGames = result["sports_content"].Game[backward:i]
			}
			forward := i + recordsNumToShow
			if forward > len(result["sports_content"].Game) {
				forward = len(result["sports_content"].Game)
			}
			nextGames := result["sports_content"].Game[i:forward]
			games := append(previousGames, nextGames...)
			for _, game := range games {
				output = append(output, []interface{}{
					p.formattedStartTime(game),
					p.formattedTeams(game),
					p.formattedScore(game),
					p.formattedPlace(game),
				})
			}
			break
		}
	}
	return output
}

package response

// GameRecord - game scheduled
type GameRecord struct {
	ID               string     `json:"id"`
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
	IsHomeTeam       string     `json:"is_home_team"`
	Outcome          string     `json:"outcome"`
}

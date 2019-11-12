package response

// PeriodTime - period statistics
type PeriodTime struct {
	GameStatus string `json:"game_status"`
	Period     string `json:"period_value"`
	Clock      string `json:"game_clock"`
}

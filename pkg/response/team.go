package response

// Team playing
type Team struct {
	ID           string `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	Nickname     string `json:"nickname"`
	URLName      string `json:"url_name"`
	Score        string `json:"score"`
}

package request

type BuyCandyRequestBody struct {
	CandyCount int    `json:"candyCount"`
	CandyType  string `json:"candyType"`
	Money      int    `json:"money"`
}

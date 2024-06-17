package request

type BuyCandyRequestBody struct {
	CandyCount int32  `json:"candyCount"`
	CandyType  string `json:"candyType"`
	Money      int32  `json:"money"`
}

package response

type ThanksResponse struct {
	Change int32  `json:"change"`
	Thanks string `json:"thanks,omitempty"`
}

type ErrorResponse struct {
	Error_ string `json:"error,omitempty"`
}

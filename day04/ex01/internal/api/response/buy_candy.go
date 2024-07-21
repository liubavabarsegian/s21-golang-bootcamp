package response

type ThanksResponse struct {
	Thanks string `json:"thanks,omitempty"`

	Change int32 `json:"change,omitempty"`
}

type ErrorResponse struct {
	Error_ string `json:"error,omitempty"`
}

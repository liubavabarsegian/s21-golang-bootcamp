package candies

import (
	"BuyCandy/internal/api/request"
	"BuyCandy/internal/api/response"
	"encoding/json"
	"errors"
	"net/http"
)

var CandiesPrices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var request request.BuyCandyRequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	change, err, statusCode := CountChange(request.CandyType, int(request.CandyCount), int(request.Money))
	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Error_: err.Error(),
		})
	} else {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response.ThanksResponse{
			Thanks: "Thank you!",
			Change: int32(change),
		})
	}
}

func CountChange(candyType string, candyCount int, money int) (change int, err error, code int) {
	price := CandiesPrices[candyType]

	if price == 0 {
		return 0, errors.New("some error in input data"), http.StatusBadRequest
	}

	if candyCount*price > money {
		return 0, errors.New("not enough money"), http.StatusPaymentRequired
	}

	return money - candyCount*price, nil, http.StatusCreated
}

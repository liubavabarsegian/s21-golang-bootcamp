package candies

import (
	"BuyCandy/internal/api/request"
	"BuyCandy/internal/api/response"
	"encoding/json"
	"errors"
	"fmt"
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

	fmt.Println("HELLO I AM HERE")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	change, err := CountChange(request.CandyType, int(request.CandyCount), int(request.Money))
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		json.NewEncoder(w).Encode(response.InlineResponse400{
			Error_: err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.InlineResponse201{
			Thanks: "Thank you!",
			Change: int32(change),
		})
	}
}

func CountChange(candyType string, candyCount int, money int) (change int, err error) {
	price := CandiesPrices[candyType]

	if candyCount*price > money {
		return 0, errors.New("not enough money")
	}

	return money - candyCount*price, nil
}

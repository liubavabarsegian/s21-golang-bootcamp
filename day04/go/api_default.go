/*
 * Candy Server
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"errors"
	"net/http"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var request BuyCandyBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	change, err := CountChange(request.CandyType, int(request.CandyCount), int(request.Money))
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		json.NewEncoder(w).Encode(InlineResponse400{
			Error_: err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(InlineResponse201{
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
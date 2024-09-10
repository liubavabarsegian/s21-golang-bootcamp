package candies

// #include "ask_cow.h"
import "C"

import (
	"OldCow/internal/api/request"
	"OldCow/internal/api/response"
	"encoding/json"
	"errors"
	"net/http"
	"unsafe"
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

	change, statusCode, err := CountChange(request.CandyType, int(request.CandyCount), int(request.Money))
	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Error_: err.Error(),
		})
	} else {
		thanks := C.CString("Thank you!")
		cowAnswer := C.ask_cow(thanks)
		defer C.free(unsafe.Pointer(cowAnswer))

		w.WriteHeader(statusCode)
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(response.ThanksResponse{
			Change: int32(change),
			Thanks: C.GoString(cowAnswer),
		})
	}
}

func CountChange(candyType string, candyCount int, money int) (change int, code int, err error) {
	price := CandiesPrices[candyType]

	if price == 0 {
		return 0, http.StatusBadRequest, errors.New("some error in input data")
	}

	if candyCount*price > money {
		return 0, http.StatusPaymentRequired, errors.New("not enough money")
	}

	return money - candyCount*price, http.StatusCreated, nil
}

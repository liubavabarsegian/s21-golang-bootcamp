package router

import (
	"OldCow/internal/api/handlers/candies"
	"net/http"
)

func SetUpRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/buy_candy", candies.BuyCandy)

	return mux
}

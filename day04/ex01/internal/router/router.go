package router

import (
	"BuyCandy/internal/api/handlers/candies"
	"fmt"
	"net/http"
)

func SetUpRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "i am protected")
	})
	mux.HandleFunc("/buy_candy", candies.BuyCandy)

	return mux
}

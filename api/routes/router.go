package routes

import (
	"net/http"
	"project/controllers"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/users", controllers.UsersHandler)

	mux.HandleFunc("/users/", controllers.UserByIDHandler)

	mux.HandleFunc("/stocks", controllers.StocksHandler)

	mux.HandleFunc("/transactions", controllers.TransactionsHandler)
}

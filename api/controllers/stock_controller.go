package controllers

import (
	"encoding/json"
	"net/http"

	"project/models"
	"project/repositories"
	"project/utils"
)

func StocksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateStock(w, r)
	case http.MethodGet:
		GetAllStocks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	err := repositories.CreateStock(stock)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusCreated, stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := repositories.GetAllStocks()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendJSON(w, http.StatusOK, stocks)
}

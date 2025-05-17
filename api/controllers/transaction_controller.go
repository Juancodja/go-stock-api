package controllers

import (
	"encoding/json"
	"net/http"

	"project/models"
	"project/repositories"
	"project/utils"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	CreateTransaction(w, r)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	exists, err := repositories.StockExist(tx.Ticker)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "database error")
		return
	}
	if !exists {
		utils.SendError(w, http.StatusBadRequest, "invalid stock ticker")
		return
	}

	id, err := repositories.CreateTransaction(tx)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tx.ID = int(id)
	utils.SendJSON(w, http.StatusCreated, tx)
}

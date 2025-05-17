package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"project/models"
	"project/repositories"
	"project/utils"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllUsers(w, r)
	case "POST":
		CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	id, err := repositories.CreateUser(user)
	if err != nil {
		http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
		return
	}

	user.ID = int(id)

	utils.SendJSON(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendJSON(w, http.StatusOK, users)
}

func UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := repositories.GetUserByID(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "user not found")
		return
	}

	utils.SendJSON(w, http.StatusOK, user)
}

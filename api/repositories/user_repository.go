package repositories

import (
	"project/config"
	"project/models"
	"database/sql"
)

func CreateUser(u models.User) (int64, error) {
	result, err := config.DB.Exec("INSERT INTO users(name, email) VALUES(?, ?)", u.Name, u.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	row := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return u, err // caller can decide how to handle "not found"
	}
	return u, err
}

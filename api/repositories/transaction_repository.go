package repositories

import (
	"project/config"
	"project/models"
)

func CreateTransaction(tx models.Transaction) (int64, error) {
	result, err := config.DB.Exec(
		`INSERT INTO transactions(user_id, ticker, type, quantity, unit_price) 
		 VALUES (?, ?, ?, ?, ?)`,
		tx.UserID, tx.Ticker, tx.Type, tx.Quantity, tx.UnitPrice,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

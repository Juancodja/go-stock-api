package repositories

import (
	"project/config"
	"project/models"
)

func CreateStock(s models.Stock) error {
	_, err := config.DB.Exec("INSERT INTO stocks(ticker, name) VALUES(?, ?)", s.Ticker, s.Name)
	return err
}

func GetAllStocks() ([]models.Stock, error) {
	rows, err := config.DB.Query("SELECT ticker, name FROM stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var s models.Stock
		if err := rows.Scan(&s.Ticker, &s.Name); err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}
	return stocks, nil
}

func StockExist(ticker string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM stocks WHERE ticker = ?)`
	err := config.DB.QueryRow(query, ticker).Scan(&exists)
	return exists, err
}
package repositories

import (
	"pricer/config"
	"pricer/models"
)

func InsertOrUpdatePrice(p models.Price) error {
	query := `
		INSERT INTO historical_prices (ticker, date, close_price)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			close_price = VALUES(close_price)
	`

	_, err := config.DB.Exec(query, p.Ticker, p.Date.Format("2006-01-02"), p.ClosePrice)
	return err
}

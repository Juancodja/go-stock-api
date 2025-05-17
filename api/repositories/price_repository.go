package repositories

import (
	"project/config"
	"project/models"
	"time"
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

func GetPriceAtDate(ticker string, date time.Time) (models.Price, error) {
	var price models.Price

	query := `
		SELECT id, ticker, date, close_price
		FROM historical_prices
		WHERE ticker = ? AND date = ?
	`

	err := config.DB.QueryRow(query, ticker, date.Format("2006-01-02")).Scan(
		&price.ID,
		&price.Ticker,
		&price.Date,
		&price.ClosePrice,
	)

	return price, err
}

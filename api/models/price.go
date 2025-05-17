package models

import "time"

type Price struct {
	ID         int       `json:"id"`
	Ticker     string    `json:"ticker"`
	Date       time.Time `json:"date"`
	ClosePrice float64   `json:"close_price"`
}

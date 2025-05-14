package models

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id" binding:"required"`
	Ticker    string  `json:"ticker" binding:"required"`
	Type      string  `json:"type" binding:"required,oneof=buy sell"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gt=0"`
}
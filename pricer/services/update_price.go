package services

import (
	"fmt"
	"log"
	"pricer/clients"
	"pricer/models"
	"pricer/repositories"
	"time"
)

func UpdatePrices() {

	stocks, err := repositories.GetAllStocks()
	if err != nil {
		log.Printf("Error retrieving stocks: %v", err)
		return
	}
	for _, s := range stocks {
		err = AddStockPrice(s)
		if err != nil {
			log.Printf("Error adding stock price for ticker %s: %v", s.Ticker, err)
			continue
		}
	}
}

func AddStockPrice(s models.Stock) error {
	response, err := clients.GetCurrentPrice(s.Ticker)
	if err != nil {
		return fmt.Errorf("error al obtener precio: %w", err)
	}

	p := response.Current
	t := TruncateDate(response.Timestamp)

	var price models.Price
	price.Ticker = s.Ticker
	price.ClosePrice = p
	price.Date = t

	err = repositories.InsertOrUpdatePrice(price)
	if err != nil {
		return fmt.Errorf("error al insertar o actualizar precio: %w", err)
	}

	return nil

}

func TruncateDate(stamp int64) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal("Failed to load location:", err)
	}
	t := time.Unix(stamp, 0).In(loc)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

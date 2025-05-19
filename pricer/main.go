package pricer

import (
	"log"
	"pricer/config"
	"pricer/services"
	"time"
)

func main() {
	config.ConnectDB()
	log.Println("Pricer started")

	ticker := time.NewTicker(2 * time.Hour)
	defer ticker.Stop()

	services.UpdatePrices()

	for range ticker.C {
		services.UpdatePrices()
	}
}

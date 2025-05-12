package main

import (
    "project/config"
    "project/routes"
)


func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}

package routes

import (
	"project/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserByID)

	// (Coming soon) Stock and transaction routes
	router.POST("/stocks", controllers.CreateStock)
	router.GET("/stocks", controllers.GetAllStocks)
	router.POST("/transactions", controllers.CreateTransaction)
	// router.GET("/users/:id/portfolio", controllers.GetPortfolio)

	return router
}

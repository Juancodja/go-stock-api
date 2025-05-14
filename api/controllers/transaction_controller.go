package controllers

import (
	"project/models"
	"project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var tx models.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := repositories.StockExist(tx.Ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid stock ticker"})
		return
	}

	id, err := repositories.CreateTransaction(tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.ID = int(id)
	c.JSON(http.StatusCreated, tx)
}

package handlers

import (
	"net/http"
	"sync"

	"github.com/fetchProject/app/models"
	"github.com/fetchProject/database"
	"github.com/fetchProject/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	receiptsMu sync.RWMutex
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a random ID for the receipt
	ID := uuid.New().String()

	// Store the receipt in memory
	receiptsMu.Lock()
	database.Receipts[ID] = receipt
	receiptsMu.Unlock()

	c.JSON(http.StatusOK, gin.H{"id": ID})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")

	// Retrieve the receipt from memory
	receiptsMu.RLock()
	receipt, exists := database.Receipts[id]
	receiptsMu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	// Calculate points based on the rules
	points, err := utils.CalculatePoints(&receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calculating points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

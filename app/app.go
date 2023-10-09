package app

import (
	"github.com/fetchProject/app/handlers"
	"github.com/gin-gonic/gin"
)

func InitHandler() *gin.Engine {
	r := gin.Default()

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetPoints)

	return r
}

package api

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/api/handler"
	"richardoctoey/interview-gorry/api/handler/transaction"
)

func StartApi(url string) {
	r := gin.Default()
	r.POST("/event/create", handler.CreateEvent)
	r.GET("/event/get_info", handler.GetEvent)
	r.POST("/event/ticket/create", handler.CreateTicket)

	r.POST("/location/create", handler.CreateLocation)

	r.POST("/transaction/purchase", transaction.TransactionPurchase)
	r.GET("/transaction/get_info", transaction.TransactionInfo)
	r.Run(url)
}

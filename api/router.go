package api

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/api/handler"
)

func StartApi(url string) {
	r := gin.Default()
	r.POST("/event/create", handler.CreateEvent)
	r.GET("/event/get_info", handler.GetEvent)
	//r.POST("/event/ticket/create", handler.TicketCreate)

	r.POST("/location/create", handler.CreateLocation)

	//r.POST("/transaction/purchase", handler.TransactionPurchase)
	//r.GET("/transaction/get_info", handler.TransactionGetInfo)
	r.Run(url)
}

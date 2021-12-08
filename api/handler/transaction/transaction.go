package transaction

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/purchase"
)

func TransactionPurchase(c *gin.Context) {
	transaction := TransactionPayload{}
	if err := c.BindJSON(&transaction); err != nil {
		common.Error(c, err)
		return
	}
	purchaseTicket := []purchase.Purchase{}
	for _, v := range transaction.Tickets {
		p := purchase.Purchase{
			CustomerName: transaction.CustomerName,
			Ticket:       v.TicketType,
			Qty:          v.Qty,
		}
		purchaseTicket = append(purchaseTicket, p)
	}
	objs, err := purchase.InsertMultipleData(purchaseTicket)
	if err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", objs)
}

func TransactionInfo(c *gin.Context) {
	listEvent, err := purchase.GetPurchases(c.Query("uuid"))
	if err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", listEvent)
}
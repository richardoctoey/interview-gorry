package handler

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/transaction"
)

func TransactionPurchase(c *gin.Context) {
	trnpayload := transaction.Transaction{}
	if err := c.BindJSON(&trnpayload); err != nil {
		common.Error(c, err)
		return
	}
	if err := trnpayload.Create(); err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", trnpayload)
}

func TransactionInfo(c *gin.Context) {
	listEvent, err := transaction.GetPurchases(c.Query("uuid"))
	if err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", listEvent)
}

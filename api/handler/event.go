package handler

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/common"
	evt "richardoctoey/interview-gorry/event"
	"richardoctoey/interview-gorry/event/ticket"
)

func CreateEvent(c *gin.Context) {
	event := evt.Event{}
	if err := c.BindJSON(&event); err != nil {
		common.Error(c, err)
		return
	}
	if err := event.Save(); err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", event)
}

func GetEvent(c *gin.Context) {
	listEvent, err := evt.GetEvents("")
	if err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", listEvent)
}

func CreateTicket(c *gin.Context) {
	tickt := ticket.Ticket{}
	if err := c.BindJSON(&tickt); err != nil {
		common.Error(c, err)
		return
	}
	if err := tickt.Save(); err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", tickt)
}
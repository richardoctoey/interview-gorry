package handler

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
)

func CreateEvent(c *gin.Context) {
	event := event.Event{}
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
	listEvent, err := event.GetEvents()
	if err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", listEvent)
}


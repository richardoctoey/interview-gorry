package handler

import (
	"github.com/gin-gonic/gin"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
)

func CreateLocation(c *gin.Context) {
	loc := event.Location{}
	if err := c.BindJSON(&loc); err != nil {
		common.Error(c, err)
		return
	}
	if err := loc.Save(); err != nil {
		common.Error(c, err)
		return
	}
	common.OK(c, "OK", loc)
}

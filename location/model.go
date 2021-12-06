package location

import (
	"errors"
	"richardoctoey/interview-gorry/common"
	"strings"
)

type Location struct {
	UUID string `gorm:"primaryKey,column:uuid" json:"uuid"`
	Name string `gorm:"column:name" json:"name"`
}

func (u Location) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Location{})
}

func (u Location) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("Name is empty")
	}
	return nil
}

package location

import (
	"errors"
	"github.com/google/uuid"
	"richardoctoey/interview-gorry/common"
	"strings"
)

type Location struct {
	UUID string `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Name string `gorm:"column:name" json:"name"`
}

func (u *Location) Save() (error) {
	if u.UUID == "" {
		if findLocationByName(u.Name) {
			return errors.New("Name already exists")
		}
		u.UUID = uuid.New().String()
		return common.GetDb().Create(&u).Error
	}

	if !findLocationByUuid(u.UUID) {
		return errors.New("Not exists")
	}
	return common.GetDb().Where("uuid = ?", u.UUID).Save(&u).Error
}

func findLocationByUuid(uuid string) bool {
	var count int64
	common.GetDb().Model(&Location{}).Where("uuid = ?", uuid).Count(&count)
	if count >= 1 {
		return true
	}
	return false
}

func findLocationByName(name string) bool {
	var count int64
	common.GetDb().Model(&Location{}).Where("name = ?", name).Count(&count)
	if count >= 1 {
		return true
	}
	return false
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

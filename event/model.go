package event

import (
	"errors"
	"richardoctoey/interview-gorry/common"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	UUID string `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Name string `gorm:"column:name" json:"name"`
	Location string `gorm:"column:location" json:"location"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime time.Time `gorm:"column:end_time" json:"end_time"`
}

func (u Event) TableName() string {
	return "event"
}

func (u *Event) Save() (error) {
	if u.UUID == "" {
		u.UUID = uuid.New().String()
		return common.GetDb().Create(&u).Error
	}

	var count int64
	common.GetDb().Model(&Event{}).Where("uuid = ?", u.UUID).Count(&count)
	if count != 1 {
		return errors.New("Not exists")
	}
	return common.GetDb().Where("uuid = ?", u.UUID).Save(&u).Error
}

func GetEvents() ([]Event, error) {
	var result []Event
	if err := common.GetDb().Model(&Event{}).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u Event) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Event{})
}

func (u Event) Validate() error {
	return nil
}


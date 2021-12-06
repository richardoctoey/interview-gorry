package event

import (
	"errors"
	"richardoctoey/interview-gorry/common"
	"time"
)

type Event struct {
	UUID string `gorm:"primaryKey,column:uuid" json:"uuid"`
	Name string `gorm:"column:name" json:"name"`
	Location string `gorm:"column:location" age:"location"`
	StartTime time.Time `gorm:"column:start_time" age:"start_time"`
	EndTime time.Time `gorm:"column:end_time" age:"end_time"`
}

func (u Event) TableName() string {
	return "event"
}

func (u *Event) Save() (error) {
	if u.UUID == "" {
		return common.GetDb().Create(&u).Error
	}

	var count int64
	common.GetDb().Model(&Event{}).Where("id = ?", u.UUID).Count(&count)
	if count != 1 {
		return errors.New("Not exists")
	}
	return common.GetDb().Where("id = ?", u.UUID).Save(&u).Error
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


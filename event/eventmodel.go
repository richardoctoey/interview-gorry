package event

import (
	"errors"
	"github.com/google/uuid"
	"richardoctoey/interview-gorry/common"
	"strings"
	"time"
)

type Event struct {
	UUID      string    `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Name      string    `gorm:"column:name" json:"name"`
	Location  string    `gorm:"column:location" json:"location"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Ticket    []Ticket  `gorm:"-" json:"tickets,omitempty"`
}

func (u Event) TableName() string {
	return "event"
}

func (u *Event) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}

	if u.UUID == "" {
		u.UUID = uuid.New().String()
		return common.GetDb().Create(&u).Error
	}

	var count int64
	common.GetDb().Model(&Event{}).Where("uuid = ?", u.UUID).Count(&count)
	if count != 1 {
		return errors.New("not exists")
	}
	return common.GetDb().Where("uuid = ?", u.UUID).Save(&u).Error
}

func (u Event) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Event{})
}

func (u Event) Validate() error {
	if strings.TrimSpace(u.Location) == "" {
		return errors.New("location cannot be empty")
	}
	if u.StartTime.IsZero() {
		return errors.New("start_time cannot be empty")
	}
	if u.EndTime.IsZero() {
		return errors.New("end_time cannot be empty")
	}
	if u.EndTime.Unix() < u.StartTime.Unix() {
		return errors.New("end_time smaller than start_time")
	}
	if isOverlapSchedule(u.Location, u.StartTime, u.EndTime) {
		return errors.New("there's another event in this period")
	}
	if !findLocationByUuid(u.Location) {
		return errors.New("no location data")
	}
	return nil
}

func isOverlapSchedule(location string, startTime time.Time, endTime time.Time) bool {
	var total int64
	common.GetDb().Model(&Event{}).Where("location = ? AND start_time < ? AND end_time > ?",
		location, endTime, startTime).Count(&total)
	if total >= 1 {
		return true
	}
	return false
}

func GetEvents(uuid string) ([]Event, error) {
	var result []Event
	find := common.GetDb().Model(&Event{})
	if uuid != "" {
		find = find.Where("uuid = ?", uuid)
	}
	if err := find.Scan(&result).Error; err != nil {
		return nil, err
	}
	if uuid != "" && len(result) > 0 {
		resultTicket, err := findByEventUuid(uuid)
		if err != nil {
			return nil, err
		}
		result[0].Ticket = resultTicket
	}
	return result, nil
}

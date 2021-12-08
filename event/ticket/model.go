package ticket

import (
	"errors"
	"github.com/google/uuid"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
	"strings"
)

type Ticket struct {
	UUID string `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Event string `gorm:"column:event" json:"event"`
	Type string `gorm:"column:type" json:"type"`
	Quota int `gorm:"column:quota" json:"quota"`
	Price float64 `gorm:"column:price" json:"price"`
}

func (u Ticket) TableName() string {
	return "ticket"
}

func (u *Ticket) Save() (error) {
	if u.UUID == "" {
		u.UUID = uuid.New().String()
		if err := u.Validate(); err != nil  {
			return err
		}
		return common.GetDb().Create(&u).Error
	}

	var count int64
	common.GetDb().Model(&Ticket{}).Where("uuid = ?", u.UUID).Count(&count)
	if count != 1 {
		return errors.New("Not exists")
	}
	if err := u.Validate(); err != nil {
		return err
	}
	return common.GetDb().Where("uuid = ?", u.UUID).Save(&u).Error
}

func eventNotExist(uuid string) bool {
	if obj, err := event.GetEvents(uuid); err == nil && len(obj) > 0 {
		return false
	}
	return true
}

func (u Ticket) Validate() error {
	if strings.TrimSpace(u.Event) == "" {
		return errors.New("event cannot be empty")
	}
	if strings.TrimSpace(u.Type) == "" {
		return errors.New("type cannot be empty")
	}
	if u.Quota == 0 {
		return errors.New("quota cannot be 0")
	}
	if eventNotExist(u.Event) {
		return errors.New("event is not exist")
	}

	//price can be 0 for now. Maybe it's a free event
	//if u.Price == 0 {
	//	return errors.New("quota cannot be 0")
	//}

	return nil
}

func (u Ticket) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Ticket{})
}
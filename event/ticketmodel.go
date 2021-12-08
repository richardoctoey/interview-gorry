package event

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"richardoctoey/interview-gorry/common"
	"strings"
)

type Ticket struct {
	UUID  string  `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Event string  `gorm:"column:event" json:"event"`
	Type  string  `gorm:"column:type" json:"type"`
	Quota int     `gorm:"column:quota" json:"quota"`
	Price float64 `gorm:"column:price" json:"price"`
}

func (u Ticket) TableName() string {
	return "ticket"
}

func (u *Ticket) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}
	if u.UUID == "" {
		u.UUID = uuid.New().String()
		return common.GetDb().Create(&u).Error
	}

	var count int64
	common.GetDb().Model(&Ticket{}).Where("uuid = ?", u.UUID).Count(&count)
	if count != 1 {
		return errors.New("not exists")
	}
	return common.GetDb().Where("uuid = ?", u.UUID).Save(&u).Error
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
	return nil
}

func (u Ticket) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Ticket{})
}

func eventNotExist(uuid string) bool {
	if obj, err := GetEvents(uuid); err == nil && len(obj) > 0 {
		return false
	}
	return true
}

func findByEventUuid(uuid string) ([]Ticket, error) {
	var result []Ticket
	if err := common.GetDb().Model(&Ticket{}).Where("event = ?", uuid).Scan(&result).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return result, nil
}

func CheckTotalQuota(uuid string) int {
	var quota int
	common.GetDb().Raw("SELECT quota FROM `ticket` WHERE uuid = ?", uuid).Scan(&quota)
	return quota
}

func CheckMultipleEvents(uuid []string) bool {
	var count int64
	common.GetDb().Raw("SELECT COUNT(*) FROM (SELECT event FROM `ticket` WHERE uuid IN (?) GROUP BY `event`) a", uuid).Scan(&count)
	if count > 1 {
		return true
	}
	return false
}

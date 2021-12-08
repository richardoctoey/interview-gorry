package purchase

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
	"richardoctoey/interview-gorry/location"
	"strings"
	"time"
)

type Purchase struct {
	UUID         string             `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	CustomerName string             `gorm:"column:customer_name" json:"customer_name"`
	Ticket       string             `gorm:"column:ticket" json:"ticket"`
	Qty          int                `gorm:"column:qty" json:"qty"`
	Event        *event.Event       `gorm:"-" json:"event,omitempty"`
	Location     *location.Location `gorm:"-" json:"location,omitempty"`
}

func (u *Purchase) Save(db *gorm.DB) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if db == nil {
		db = common.GetDb()
	}
	if u.UUID == "" {
		u.UUID = uuid.New().String()
		return common.GetDb().Create(&u).Error
	}
	return db.Where("uuid = ?", u.UUID).Save(&u).Error
}

func (u Purchase) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Purchase{})
}

func (u Purchase) Validate() error {
	if strings.TrimSpace(u.CustomerName) == "" {
		return errors.New("customer_name is empty")
	}
	if strings.TrimSpace(u.Ticket) == "" {
		return errors.New("ticket is empty")
	}
	if u.Qty == 0 {
		return errors.New("qty cannot be 0")
	}
	if err := ticketQuotaStillAvailable(u.Ticket, u.Qty); err != nil {
		return err
	}
	return nil
}

func ticketQuotaStillAvailable(uuid string, bought int) error {
	var totalBuy int
	common.GetDb().Raw("SELECT SUM(qty) FROM `purchases` WHERE ticket = ?", uuid).Scan(&totalBuy)
	totalBuy += bought
	if totalBuy > event.CheckTotalQuota(uuid) {
		return errors.New("quota is reached!")
	}
	return nil
}

func GetPurchases(uuid string) ([]Purchase, error) {
	var result []map[string]interface{}
	resultMapped := []Purchase{}
	sql := `SELECT p.*, e.uuid event_uuid, e.name event_name, e.location event_location, 
e.start_time event_start_time, e.end_time event_end_time, l.uuid location_uuid, l.name location_name 
FROM purchases p INNER JOIN ticket t ON t.uuid = p.ticket 
INNER JOIN event e ON t.event = e.uuid 
INNER JOIN locations l ON l.uuid = e.location`
	if uuid != "" {
		sql += ` WHERE p.uuid=?`
	}
	find := common.GetDb().Raw(sql, uuid)
	if err := find.Scan(&result).Error; err != nil {
		return nil, err
	}
	for _, v := range result {
		r := Purchase{
			UUID:         v["uuid"].(string),
			CustomerName: v["customer_name"].(string),
			Ticket:       v["ticket"].(string),
			Qty:          int(v["qty"].(int64)),
			Event: &event.Event{
				UUID:      v["event_uuid"].(string),
				Name:      v["event_name"].(string),
				Location:  v["event_location"].(string),
				StartTime: v["event_start_time"].(time.Time),
				EndTime:   v["event_end_time"].(time.Time),
			},
			Location: &location.Location{
				UUID: v["location_uuid"].(string),
				Name: v["location_name"].(string),
			},
		}
		resultMapped = append(resultMapped, r)
	}
	return resultMapped, nil
}

func InsertMultipleData(datas []Purchase) ([]Purchase, error) {
	db := common.GetDb().Begin()
	ticketUuid := []string{}
	for i, _ := range datas {
		if err := datas[i].Save(db); err != nil {
			db.Rollback()
			return datas, err
		}
		ticketUuid = append(ticketUuid, datas[i].Ticket)
	}
	if event.CheckMultipleEvents(ticketUuid) {
		db.Rollback()
		return nil, errors.New("cannot purchase multiple event in one transaction")
	}
	db.Commit()
	return datas, nil
}

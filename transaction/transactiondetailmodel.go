package transaction

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
	"strings"
)

type TransactionDetail struct {
	UUID        string          `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	Transaction string          `gorm:"column:transaction" json:"transaction"`
	Ticket      string          `gorm:"column:ticket" json:"ticket"`
	Qty         int             `gorm:"column:qty" json:"qty"`
	Event       *event.Event    `gorm:"-" json:"event,omitempty"`
	Location    *event.Location `gorm:"-" json:"location,omitempty"`
}

func (u TransactionDetail) TableName() string {
	return "transaction_detail"
}

func (u TransactionDetail) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&TransactionDetail{})
}

func (u TransactionDetail) Validate() error {
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

func (u *TransactionDetail) Save(db *gorm.DB) error {
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

func ticketQuotaStillAvailable(uuid string, bought int) error {
	var totalBuy int
	common.GetDb().Raw("SELECT COALESCE(SUM(qty),0) FROM `transaction_detail` WHERE ticket = ?", uuid).Scan(&totalBuy)
	totalBuy += bought
	if totalBuy > event.CheckTotalQuota(uuid) {
		return errors.New("quota is reached!")
	}
	return nil
}

func saveNewDetail(transactionUuid string, detail []TransactionDetail, db *gorm.DB) error {
	ticketUuid := []string{}
	for i, _ := range detail {
		detail[i].Transaction = transactionUuid
		if err := detail[i].Save(db); err != nil {
			return err
		}
		ticketUuid = append(ticketUuid, detail[i].Ticket)
	}
	if event.CheckMultipleEvents(ticketUuid) {
		return errors.New("cannot transaction multiple event in one transaction")
	}
	return nil
}

func getTransactionDetailByTransactionUuid(uuid string) ([]TransactionDetail, error) {
	var result []TransactionDetail
	err := common.GetDb().Model(&TransactionDetail{}).Where("transaction = ?", uuid).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

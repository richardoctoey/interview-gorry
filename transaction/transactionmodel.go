package transaction

import (
	"errors"
	"github.com/google/uuid"
	"richardoctoey/interview-gorry/common"
	"strings"
)

type Transaction struct {
	UUID         string              `gorm:"primaryKey,column:uuid" json:"uuid,omitempty"`
	CustomerName string              `gorm:"column:customer_name" json:"customer_name"`
	Ticket       string              `gorm:"-" json:"ticket,omitempty"`
	Qty          int                 `gorm:"-" json:"qty,omitempty"`
	Detail       []TransactionDetail `gorm:"-" json:"detail,omitempty"`
}

func (u *Transaction) Create() error {
	if err := u.Validate(); err != nil {
		return err
	}
	u.UUID = uuid.New().String()
	db := common.GetDb().Begin()
	err := db.Create(&u).Error
	if err != nil {
		db.Rollback()
		return err
	}
	if err := saveNewDetail(u.UUID, u.Detail, db); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func (u Transaction) AutoMigrate() error {
	return common.GetDb().AutoMigrate(&Transaction{})
}

func (u Transaction) Validate() error {
	if strings.TrimSpace(u.CustomerName) == "" {
		return errors.New("customer_name is empty")
	}
	return nil
}

func GetPurchases(uuid string) ([]Transaction, error) {
	result := []Transaction{}
	find := common.GetDb().Model(&Transaction{})
	if uuid != "" {
		find = find.Where("uuid = ?", uuid)
	}
	if err := find.Scan(&result).Error; err != nil {
		return nil, err
	}
	for i, _ := range result {
		if obj, err := getTransactionDetailByTransactionUuid(result[i].UUID); err != nil {
			return nil, err
		} else {
			result[i].Detail = obj
		}
	}
	return result, nil
}

package repository

import (
	"accounting/datamodals"
	"accounting/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IBillRepository interface {
	InsertBill(bill *datamodals.Bill, categoryId int) error
}

func NewBillRepository() IBillRepository {
	db, err := util.Connection()
	if err != nil {
		logrus.Error(err)
	}

	return &BillRepository{db:db}
}

type BillRepository struct {
	db *gorm.DB
}

func (b BillRepository) InsertBill(bill *datamodals.Bill, categoryId int) error {
	return b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bill).Error; err != nil {
			return err
		}

		if err := tx.Create(&datamodals.BillCategory{BillID:bill.ID, CategoryID:categoryId}).Error; err != nil {
			return err
		}

		return nil
	})
}


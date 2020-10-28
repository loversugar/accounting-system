package repository

import (
	"accounting/constants"
	"accounting/datamodals"
	"accounting/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IBillRepository interface {
	InsertBill(bill *datamodals.Bill, category *datamodals.BillCategory) error
	SelectBillsByDate(userId int, startDate, endDate string) []*datamodals.BillCategoryResult
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

func (b BillRepository) SelectBillsByDate(userId int, startDate, endDate string) []*datamodals.BillCategoryResult {
	var results []*datamodals.BillCategoryResult
	b.db.Debug().
		Table(constants.Bill).
		Select(constants.Bill+".id as bill_id, " +
			constants.Bill + ".consumption as consumption, " +
			constants.Bill + ".note as note, " +
			constants.Bill +".selected_time as selected_time, " +
			constants.BillCategory +".category_name as category_name, " +
			constants.BillCategory + ".category_url as category_url").
		Joins(
			"join " + constants.BillCategory +
				" on " + constants.Bill+  ".id = " + constants.BillCategory + ".bill_id").
		Where(
			constants.Bill + ".user_id = ? and " +
				constants.Bill + ".selected_time between ? and ?",
				userId, startDate, endDate).Scan(&results)

	return results
}

func (b BillRepository) InsertBill(bill *datamodals.Bill, category *datamodals.BillCategory) error {
	return b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bill).Error; err != nil {
			return err
		}

		category.BillID = bill.ID

		if err := tx.Create(&category).Error; err != nil {
			return err
		}

		return nil
	})
}


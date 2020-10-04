package datamodals

import "time"

type Bill struct {
	ID int
	UserId string `gorm:"column:user_id"`
	Consumption float32 `gorm:"column:consumption"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type BillCategory struct {
	BillID int `gorm:"primaryKey;column:bill_id"`
	CategoryID int `gorm:"primaryKey;column:category_id"`
	CategoryName string `gorm:"column:category_name"`
	CategoryUrl string `gorm:"column:category_url"`
}

type Category struct {
	ID int
	CategoryName string `gorm:"column:category_name"`
	CategoryUrl string `gorm:"column:category_url"`
	UserId int `gorm:"column:user_id"'`
	IsPrivate bool `gorm:"column:is_private"`
	Deleted bool `gorm:"column:deleted"`
}

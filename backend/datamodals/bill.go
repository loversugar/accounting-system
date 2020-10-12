package datamodals

import "time"

type Bill struct {
	ID int
	UserId string `gorm:"column:user_id"`
	Consumption float32 `gorm:"column:consumption"`
	Note string `gorm:"column:note"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type BillCategory struct {
	BillID int `gorm:"primaryKey;column:bill_id"`
	CategoryID int `gorm:"primaryKey;column:category_id"`
	CategoryName string `gorm:"column:category_name"`
	CategoryUrl string `gorm:"column:category_url"`
}

type Category struct {
	ID int `json:"id"`
	CategoryName string `json:"categoryName" gorm:"column:category_name"`
	CategoryUrl string `json:"categoryUrl" gorm:"column:category_url"`
	UserId int `json:"userId" gorm:"column:user_id"'`
	IsPrivate bool `json:"-" gorm:"column:is_private"`
	Deleted bool `json:"-" gorm:"column:deleted"`
}

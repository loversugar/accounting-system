package datamodals

import "time"

type Bill struct {
	ID int
	OpenId string `gorm:"column:open_id"`
	Username string `gorm:"column:username"`
	Consumption float32 `gorm:"column:consumption"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type BillCategory struct {
	BillID int `gorm:"primaryKey;column:bill_id"`
	CategoryID int `gorm:"primaryKey;column:category_id"`
}

type Category struct {
	ID int
	CategoryName string `gorm:"column:category_name"`
	CategoryUrl string `gorm:"column:category_url"`
	UserId int `gorm:"column:user_id"'`
}

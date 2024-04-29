package model

import "gorm.io/gorm"

// 活动表 销售要约
type Goods struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30)"`
	Price float32 `gorm:"type:decimal(10,2)"`
	Stock int     `gorm:"type:int(11)"`
}

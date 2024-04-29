package model

import "gorm.io/gorm"

// 活动表 销售要约
type ActivityGoods struct {
	gorm.Model
	ActivityID int `gorm:"type:int(11);column:activity_id;not null"`
	GoodsID    int `gorm:"type:int(11);column:goods_id;not null"`
}

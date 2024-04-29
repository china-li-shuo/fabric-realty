package model

import "gorm.io/gorm"

// 活动表 销售要约
type Activity struct {
	gorm.Model
	Name      string `gorm:"type:varchar(30);column:name;"`
	StartTime int    `gorm:"type:int(10);column:start_time;"`
	EndTime   int    `gorm:"type:int(10);column:end_time;"`
}

package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/fabirc_fang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Activity{}, &Goods{}, &ActivityGoods{}, &Order{})
}

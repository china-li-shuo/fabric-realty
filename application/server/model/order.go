package model

import "gorm.io/gorm"

// 订单表
type Order struct {
	gorm.Model
	OrderSn    string  `gorm:"type:varchar(30);column:order_sn;not null;"`
	UserID     int     `gorm:"type:int(11);column:user_id;not null;"`
	GoodsID    int     `gorm:"type:int(11);column:goods_id;not null;"`
	ActivityID int     `gorm:"type:int(11);column:activity_id;not null;"`
	Price      float32 `gorm:"type:decimal(10,2);column:price;not null;"`
	Stock      int     `gorm:"type:int(11);column:stock;not null;"`
	PayType    int     `gorm:"type:tinyint(1);column:pay_type;not null;comment:1支付宝 2微信"`
	Status     int8    `gorm:"type:tinyint(1);column:status;not null;comment:0待支付 1已支付 2交易关闭 3退款 "`
}

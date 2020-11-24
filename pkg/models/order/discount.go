package order

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Discounts struct {
	DiscountList []Discount // 折扣列表
}

type Discount struct {
	OrderId          uint64       `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`               // 订单ID
	SerialId         uint8        `gorm:"column:serialId;primary_key;auto_increment:false;default:'0';comment:'计算序列号'"` // 计算序列号
	DiscountType     DiscountType `gorm:"column:discountType;default:'0';comment:'折扣类型（折扣类型，如满减券、代金券）'"`                // 折扣类型（折扣类型，如满减券、代金券）
	DiscountPrice    uint64       `gorm:"column:discountPrice;type:bigint unsigned;default:'0';comment:'折扣金额（单位：分）'"`   // 折扣金额（单位：分）
	DiscountObjectId string       `gorm:"column:discountObjectId;type:varchar(40);default:'';comment:'相关折扣对象ID'"`       // 相关折扣对象ID
	databases.Time
}

func (Discount) TableName() string {
	return "order_discount"
}

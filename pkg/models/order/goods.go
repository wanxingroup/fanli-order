package order

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Goods struct {
	OrderId  uint64 `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`       // 订单ID
	SkuId    uint64 `gorm:"column:skuId;primary_key;auto_increment:false;comment:'SKU ID'"`       // SKU ID
	GoodsId  uint64 `gorm:"column:goodsId;default:'0';comment:'商品ID'"`                            // 商品ID
	Price    uint64 `gorm:"column:price;type:bigint unsigned;default:'0';comment:'商品单价（单位：分）'"`   // 商品单价（单位：分）
	VipPrice uint64 `gorm:"column:vipPrice;type:bigint unsigned;default:'0';comment:'会员价（单位：分）'"` // 会员价（单位：分）
	Point    uint64 `gorm:"column:point;type:bigint unsigned;default:'0';comment:'积分'"`           // 积分（单位：分）
	Count    uint32 `gorm:"column:count;type:int unsigned;default:'0';comment:'商品数量'"`            // 商品数量
	databases.Time
}

func (Goods) TableName() string {
	return "order_goods"
}

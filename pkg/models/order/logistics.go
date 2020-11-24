package order

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Logistics struct {
	OrderId          uint64 `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`       // 订单ID
	PackageId        uint64 `gorm:"column:packageId;primary_key;auto_increment:false;comment:'包裹ID'"`     // 包裹ID
	LogisticsCompany string `gorm:"column:logisticsCompany;type:varchar(64);default:'';comment:'物流公司名称'"` // 物流公司名称
	ExpressNumber    string `gorm:"column:expressNumber;type:varchar(64);default:'';comment:'快递单号'"`      // 快递单号
	databases.Time
}

func (Logistics) TableName() string {
	return "order_logistics"
}

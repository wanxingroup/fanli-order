package order

import (
	"github.com/jinzhu/gorm"
)

type StatusModificationLog struct {
	gorm.Model
	OrderId           uint64 `gorm:"column:orderId;index:orderId;default:'0';comment:'订单ID'"` // 订单ID
	DestinationStatus Status `gorm:"column:destinationStatus;default:'0';comment:'目标订单状态'"`   // 目标订单状态
}

func (StatusModificationLog) TableName() string {
	return "order_status_modification_log"
}

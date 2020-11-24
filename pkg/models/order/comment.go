package order

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Comment struct {
	OrderId   uint64 `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"` // 订单ID
	Commented bool   `gorm:"column:commented;default:'0';comment:'是否已经评论'"`                  // 是否已经评论
	databases.Time
}

func (Comment) TableName() string {
	return "order_comment"
}

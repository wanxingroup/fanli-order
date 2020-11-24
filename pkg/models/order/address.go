package order

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Address struct {
	OrderId      uint64 `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`  // 订单ID
	Province     string `gorm:"column:province;type:varchar(32);default:'';comment:'省份'"`        // 省份
	City         string `gorm:"column:city;type:varchar(32);default:'';comment:'城市'"`            // 城市
	District     string `gorm:"column:district;type:varchar(32);default:'';comment:'辖区'"`        // 辖区
	Street       string `gorm:"column:street;type:varchar(32);default:'';comment:'街道'"`          // 街道
	Address      string `gorm:"column:address;type:varchar(128);default:'';comment:'具体地址'"`      // 具体地址
	ReceiverName string `gorm:"column:receiverName;type:varchar(32);default:'';comment:'收件人名称'"` // 收件人名称
	Tel          string `gorm:"column:tel;type:varchar(32);default:'';comment:'联系电话'"`           // 联系电话
	databases.Time
}

func (Address) TableName() string {
	return "order_address"
}

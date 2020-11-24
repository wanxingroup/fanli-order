package order

import (
	"time"

	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Payment struct {
	OrderId        uint64        `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`                              // 订单ID
	TransactionId  string        `gorm:"column:transactionId;type:varchar(40);primary_key;auto_increment:false;comment:'支付网关支付事务ID'"` // 支付网关支付事务ID
	PaidPrice      uint64        `gorm:"column:paidPrice;type:bigint unsigned;not null;default:'0';comment:'实付金额（单位：分）'"`             // 实付金额（单位：分）
	PaymentChannel string        `gorm:"column:paymentChannel;type:varchar(30);default:'';comment:'支付渠道（富友、汇付）'"`                     // 支付渠道（富友、汇付）
	PaymentMode    string        `gorm:"column:paymentMode;type:varchar(30);default:'';comment:'支付方式（H5、APP、公众号、主被扫码）'"`              // 支付方式（H5、APP、公众号、主被扫码）
	PaymentProduct string        `gorm:"column:status;type:varchar(30);default:'';comment:'支付产品（微信、支付宝）'"`                            // 支付产品（微信、支付宝）
	PaidTime       *time.Time    `gorm:"column:paidTime;comment:'支付时间'"`                                                              // 支付时间
	Status         PaymentStatus `gorm:"column:paymentStatus;type:tinyint unsigned;not null;comment:'支付流水状态'"`                        // 支付流水状态
	databases.Time
}

func (Payment) TableName() string {
	return "order_payment"
}

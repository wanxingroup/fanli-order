package order

import (
	"time"

	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Information struct {
	OrderId           uint64                   `gorm:"column:orderId;primary_key;auto_increment:false;comment:'订单ID'"`                   // 订单ID
	OrderType         Type                     `gorm:"column:orderType;default:'1';index:orderType_idx;comment:'订单类型'"`                  // 订单类型
	MerchantId        uint64                   `gorm:"column:merchantId;default:'0';comment:'商家ID'"`                                     // 商家ID
	ShopId            uint64                   `gorm:"column:shopId;default:'0';comment:'店铺ID'"`                                         // 店铺ID
	UserId            uint64                   `gorm:"column:userId;default:'0';comment:'用户ID'"`                                         // 用户ID
	Status            Status                   `gorm:"column:status;default:'0';index:status_idx;comment:'订单状态'"`                        // 订单状态
	RefundStatus      RefundStatus             `gorm:"column:refundStatus;default:'0';index:refundStatus_idx;comment:'退款状态'"`            // 退款状态
	GoodsTotalAmount  uint64                   `gorm:"column:goodsTotalAmount;type:bigint unsigned;default:'0';comment:'订单商品总金额（单位：分）'"` // 订单商品总金额（单位：分）
	DiscountPrice     uint64                   `gorm:"column:discountPrice;type:bigint unsigned;default:'0';comment:'订单折扣金额（单位：分）'"`     // 订单折扣金额（单位：分）
	Payable           uint64                   `gorm:"column:payable;type:bigint unsigned;default:'0';comment:'订单应付金额（单位：分）'"`           // 订单应付金额（单位：分）
	Freight           uint64                   `gorm:"column:freight;type:bigint unsigned;default:'0';comment:'运费（单位：分）'"`               // 运费（单位：分）
	Point             uint64                   `gorm:"column:point;type:bigint unsigned;default:'0';comment:'订单应付积分'"`                   // 订单应付积分
	IsVip             bool                     `gorm:"column:isVip;type:bool;default:'0';comment:'是否VIP购买'"`                             // 是否VIP购买
	Remark            string                   `gorm:"column:remark;type:varchar(200);default:'';comment:'备注'"`
	PaidTime          *time.Time               `gorm:"column:paidTime;comment:'支付时间'"`
	LogisticsPackages []*Logistics             `gorm:"foreignkey:orderId;comment:'物流包裹'"`     // 物流包裹
	GoodsList         []*Goods                 `gorm:"foreignkey:orderId;comment:'商品列表'"`     // 商品列表
	Discounts         []*Discount              `gorm:"foreignkey:orderId;comment:'折扣信息'"`     // 折扣信息
	Address           *Address                 `gorm:"foreignkey:orderId;comment:'收货地址信息'"`   // 收货地址信息
	Payments          []*Payment               `gorm:"foreignkey:orderId;comment:'支付信息'"`     // 支付信息
	ModificationLogs  []*StatusModificationLog `gorm:"foreignkey:orderId;comment:'订单状态修改日志'"` // 订单状态修改日志
	Comment           *Comment                 `gorm:"foreignkey:orderId;comment:'订单评论信息'"`   // 订单评论信息
	databases.Time
}

func (Information) TableName() string {
	return "order_information"
}

package query

import (
	"time"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type QueryOrders struct {
	Status           order.Status
	RefundStatusList []order.RefundStatus
	OrderType        order.Type
	UserId           uint64
	UserIds          []uint64
	ShopId           uint64
	MerchantId       uint64
	OderId           uint64
	OrderIds         []uint64
	PageSize         uint64
	Page             uint64
	CreateTime       *TimeRange
	PaidTime         *TimeRange
	LastOrderId      uint64
}

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

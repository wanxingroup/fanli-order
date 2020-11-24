package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Interface interface {
	// 是否可以变更为指定状态
	CanChangeTo(status order.Status) bool
	// 输出当前状态值
	GetStatus() order.Status
	// 处理状态变更前的事务
	OnPrevChange(order *order.Information) error
	// 处理状态变更后的事务
	OnChanged(order *order.Information) error
}

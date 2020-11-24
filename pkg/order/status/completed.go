package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Completed struct {
}

func (_ Completed) CanChangeTo(status order.Status) bool {

	// 订单完成是一个最终状态
	return false
}

func (_ Completed) GetStatus() order.Status {

	return order.StatusCompleted
}

func (_ Completed) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ Completed) OnChanged(order *order.Information) error {

	return nil
}

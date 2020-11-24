package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Closed struct {
}

func (_ Closed) CanChangeTo(status order.Status) bool {

	// 订单关闭是一个最终状态
	return false
}

func (_ Closed) GetStatus() order.Status {

	return order.StatusClosed
}

func (_ Closed) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ Closed) OnChanged(order *order.Information) error {

	return nil
}

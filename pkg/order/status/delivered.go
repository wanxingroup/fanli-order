package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Delivered struct {
}

func (_ Delivered) CanChangeTo(status order.Status) bool {

	switch status {
	case order.StatusDelivered, order.StatusReceived, order.StatusClosed:
		return true
	}

	return false
}

func (_ Delivered) GetStatus() order.Status {

	return order.StatusDelivered
}

func (_ Delivered) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ Delivered) OnChanged(order *order.Information) error {

	return nil
}

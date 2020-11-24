package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Paid struct {
}

func (_ Paid) CanChangeTo(status order.Status) bool {

	switch status {
	case order.StatusDelivered, order.StatusClosed:
		return true
	}

	return false
}

func (_ Paid) GetStatus() order.Status {

	return order.StatusPaid
}

func (_ Paid) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ Paid) OnChanged(order *order.Information) error {

	return nil
}

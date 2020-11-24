package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Received struct {
}

func (_ Received) CanChangeTo(status order.Status) bool {

	switch status {
	case order.StatusClosed, order.StatusCompleted:
		return true
	}

	return false
}

func (_ Received) GetStatus() order.Status {

	return order.StatusReceived
}

func (_ Received) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ Received) OnChanged(order *order.Information) error {

	return nil
}

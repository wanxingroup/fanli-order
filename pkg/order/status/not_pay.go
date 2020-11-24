package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type NotPay struct {
}

func (_ NotPay) CanChangeTo(status order.Status) bool {

	switch status {
	case order.StatusCancel, order.StatusPaid:
		return true
	}

	return false
}

func (_ NotPay) GetStatus() order.Status {

	return order.StatusNotPay
}

func (_ NotPay) OnPrevChange(order *order.Information) error {

	return nil
}

func (_ NotPay) OnChanged(order *order.Information) error {

	return nil
}

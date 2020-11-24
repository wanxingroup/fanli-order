package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func newStatus(status order.Status) Interface {

	switch status {
	case order.StatusNotPay:
		return &NotPay{}
	case order.StatusPaid:
		return &Paid{}
	case order.StatusDelivered:
		return &Delivered{}
	case order.StatusReceived:
		return &Received{}
	case order.StatusClosed:
		return &Closed{}
	case order.StatusCompleted:
		return &Completed{}
	case order.StatusCancel:
		return &Cancel{}
	}
	return &null{}
}

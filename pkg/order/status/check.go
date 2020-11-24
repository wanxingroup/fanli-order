package status

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func CanChangeTo(status order.Status, destinationStatus order.Status) bool {

	return newStatus(status).CanChangeTo(destinationStatus)
}

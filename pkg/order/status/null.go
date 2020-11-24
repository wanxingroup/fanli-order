package status

import (
	"errors"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

const notExistStatus order.Status = 0

var nullObjectError = errors.New("null object")

type null struct {
}

func (_ null) CanChangeTo(status order.Status) bool {

	return false
}

func (_ null) GetStatus() order.Status {

	return notExistStatus
}

func (_ null) OnPrevChange(order *order.Information) error {

	return nullObjectError
}

func (_ null) OnChanged(order *order.Information) error {

	return nullObjectError
}

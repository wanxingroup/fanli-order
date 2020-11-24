package cancel

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/status"
)

var ErrorOrderNotExist = errors.New("order does not exist")
var ErrorOrderStatusCannotCancel = errors.New("order status cannot be cancelled")

func Cancel(orderId uint64) error {
	information, err := query.GetOrder(orderId)
	if err != nil {
		log.WithFields(log.Fields{
			"orderId": orderId,
		}).Error(err)
		return err
	}
	if information == nil {
		return ErrorOrderNotExist
	}
	if information.Status != order.StatusNotPay {
		return ErrorOrderStatusCannotCancel
	}

	err = status.ChangeStatus(information, order.StatusCancel)
	if err != nil {
		return err
	}
	return nil
}

func AutoCancelTime(orderCreateTime time.Time) time.Time {

	return orderCreateTime.Add(constants.AutoCancelOrderSeconds * time.Second)
}

func BeforeTimeToCreateOrderNeedCancel() time.Time {
	return time.Now().Add(-constants.AutoCancelOrderSeconds * time.Second)
}

func IsOrderNotExistError(err error) bool {
	return err == ErrorOrderNotExist
}

func IsOrderStatusCannotCancel(err error) bool {
	return err == ErrorOrderStatusCannotCancel
}

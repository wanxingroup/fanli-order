package logistics

import (
	"errors"

	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/status"
)

var errorStatusCannotReceived = errors.New("status not satisfy received")

func Received(data *order.Information) (err error) {
	if !status.CanChangeTo(data.Status, order.StatusReceived) {
		logrus.WithField("order", data).Warn("order status not satisfy to received")
		return errorStatusCannotReceived
	}
	err = status.ChangeStatus(data, order.StatusReceived)
	if err != nil {
		logrus.WithError(err).WithField("order", data).Error("change status error")
		return err
	}

	return nil
}

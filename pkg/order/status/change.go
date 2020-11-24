package status

import (
	"errors"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

var ErrorCannotChangeToSpecifyStatus = errors.New("it can't change the status to specify status")
var ErrorStatusWasChanged = errors.New("order status was changed by other request")

func ChangeStatus(data *order.Information, status order.Status) error {

	currentStatus := newStatus(data.Status)

	if !currentStatus.CanChangeTo(status) {

		logrus.WithField("order", data).
			WithField("targetStatus", status).
			Error(ErrorCannotChangeToSpecifyStatus)
		return ErrorCannotChangeToSpecifyStatus
	}

	statusLogical := newStatus(status)
	if err := statusLogical.OnPrevChange(data); err != nil {

		logrus.WithField("order", data).
			WithField("targetStatus", status).
			Error(err)
		return err
	}

	if err := setStatus(data, status); err != nil {

		logrus.WithField("order", data).
			WithField("targetStatus", status).
			Error(err)
		return err
	}

	if err := statusLogical.OnChanged(data); err != nil {

		logrus.WithField("order", data).
			WithField("targetStatus", status).
			Error(err)

		// 需要先回滚状态
		if rollbackError := setStatus(data, currentStatus.GetStatus()); rollbackError != nil {

			// 如果Rollback错误的话，只记录状态，后面可以逐步开发通知和错误类型收集机制，并安排解决
			logrus.WithField("order", data).
				WithField("targetStatus", status).
				WithField("rollbackStatus", currentStatus.GetStatus()).
				WithField("action", "rollback").
				Error(rollbackError)
		}

		return err
	}

	if err := addModificationLog(data); err != nil {

		logrus.WithField("order", data).
			WithField("targetStatus", status).
			WithField("action", "add modification log").
			Error(err)
	}

	return nil
}

func setStatus(data *order.Information, status order.Status) error {

	err := database.GetDB(constants.DatabaseConfigKey).
		Model(data).
		Where(order.Information{Status: data.Status}).
		Select("status").
		Updates(order.Information{Status: status}).
		Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrorStatusWasChanged
		}

		return err
	}

	data.Status = status

	return nil
}

func addModificationLog(data *order.Information) (err error) {

	modificationLog := &order.StatusModificationLog{
		OrderId:           data.OrderId,
		DestinationStatus: data.Status,
	}

	err = database.GetDB(constants.DatabaseConfigKey).Create(modificationLog).Error

	if data.ModificationLogs == nil {
		return
	}

	data.ModificationLogs = append(data.ModificationLogs, modificationLog)

	return
}

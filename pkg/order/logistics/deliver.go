package logistics

import (
	"errors"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idCreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/status"
)

var errorStatusCannotDeliver = errors.New("status not satisfy deliver")

func Deliver(data *order.Information, logistics *order.Logistics) (resultData *order.Logistics, err error) {

	if !status.CanChangeTo(data.Status, order.StatusDelivered) {

		logrus.WithField("order", data).Warn("order status not satisfy to deliver")
		return nil, errorStatusCannotDeliver
	}

	logisticsData := *logistics
	logisticsData.OrderId = data.OrderId
	logisticsData.PackageId = idCreator.NextID()

	transaction := database.GetDB(constants.DatabaseConfigKey).Begin()
	defer func() {
		if err != nil {
			errRollback := transaction.Rollback().Error
			if errRollback != nil {

				// 如果回滚失败，记录错误
				logrus.WithField("logistics", logisticsData).
					WithError(errRollback).
					Error("rollback logistics record error")
			}
		} else {
			transaction.Commit()
		}
	}()

	err = createLogisticsRecord(transaction, &logisticsData)
	if err != nil {
		logrus.WithField("error", err).Error("create package record error")
		return nil, err
	}

	data.LogisticsPackages = append(data.LogisticsPackages, &logisticsData)

	err = status.ChangeStatus(data, order.StatusDelivered)
	if err != nil {

		logrus.WithError(err).WithField("order", data).Error("change status error")
		return nil, err
	}

	return &logisticsData, nil
}

func createLogisticsRecord(transaction *gorm.DB, data *order.Logistics) (err error) {

	return transaction.Model(&order.Logistics{}).Create(data).Error
}

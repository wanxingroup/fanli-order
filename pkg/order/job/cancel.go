package job

import (
	"strconv"
	"strings"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/cancel"
)

const PageSize int = 20

func CancelOrdersJob() {
	informationCancelCron := order.InformationCancelCron{
		CronId: "CancelCron",
	}
	err := createCancelCronRecord(&informationCancelCron)
	if err == nil {
		cancelOrders()
		return
	}
	cancelCron, err := queryCancelCronRecord()
	if err == gorm.ErrRecordNotFound {
		return
	}
	if err != nil {
		logrus.WithError(err).Error("get cron record error")
		return
	}
	if cancelCron == nil {
		return
	}
	if num := updateCancelCron(cancelCron); num == 0 {
		return
	}
	cancelOrders()
}

func cancelOrders() {
	count, err := getOrdersNum(uint8(order.StatusNotPay))
	if err != nil {
		logrus.WithField("GetOrdersNum", count).
			WithField("error", err).
			Info("自动取消订单---查询未支付订单失败")
		return
	}
	if count == 0 {
		logrus.Info("自动取消订单---未支付的订单数量为0")
		return
	}
	var pageNum int = 1
	if count%PageSize == 0 {
		pageNum = count / PageSize
	} else {
		pageNum = (count / PageSize) + 1
	}
	var pageNo int = 1
	for pageNo = 1; pageNo <= pageNum; pageNo++ {
		orderList, err := orders(uint8(order.StatusNotPay), PageSize, pageNo)
		go func() {
			for _, information := range orderList {
				err = cancel.Cancel(information.OrderId)
				if err != nil {
					logrus.WithError(err).WithField("orderId", information.OrderId).Error("自动取消订单---取消订单失败")
				}
			}
		}()
	}
}

func getOrdersNum(status uint8) (count int, err error) {
	err = database.GetDB(constants.DatabaseConfigKey).Model(&order.Information{}).
		Where(getCondition(status, constants.AutoCancelOrderSeconds)).Count(&count).Error
	return count, err
}

func orders(status uint8, pageSize, pageNum int) (orderList []order.Information, err error) {
	err = database.GetDB(constants.DatabaseConfigKey).Model(&order.Information{}).
		Where(getCondition(status, constants.AutoCancelOrderSeconds)).
		Offset((pageNum - 1) * pageSize).
		Find(&orderList).Error
	return orderList, err
}

func createCancelCronRecord(data *order.InformationCancelCron) (err error) {
	return database.GetDB(constants.DatabaseConfigKey).Model(&order.InformationCancelCron{}).Create(data).Error
}

func queryCancelCronRecord() (*order.InformationCancelCron, error) {
	cancelCron := new(order.InformationCancelCron)
	err := database.GetDB(constants.DatabaseConfigKey).Model(&order.InformationCancelCron{}).
		Where("createdAt <= DATE_ADD(NOW(), INTERVAL - 1 MINUTE)").
		First(&cancelCron).Error
	return cancelCron, err
}

func updateCancelCron(data *order.InformationCancelCron) (num int64) {

	num = database.GetDB(constants.DatabaseConfigKey).Model(data).Where(
		order.InformationCancelCron{CronId: data.CronId,
			CreatedAt: data.CreatedAt}).Update(order.InformationCancelCron{
		CreatedAt: time.Now()}).RowsAffected
	return num
}

func getCondition(status uint8, cancelPaySeconds int) string {
	var condition strings.Builder
	condition.WriteString("status = " + strconv.Itoa(int(status)))
	condition.WriteString(" AND createdAt <= DATE_ADD(NOW(), INTERVAL - " + strconv.Itoa(cancelPaySeconds) + " SECOND)")
	return condition.String()
}

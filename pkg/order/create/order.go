package create

import (
	"context"
	"time"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	idCreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/create/operation"
)

// 新建订单结构数据
func NewOrder(orderId uint64, options ...Option) *order.Information {

	if orderId <= 0 {
		orderId = idCreator.NextID()
	}

	orderData := &order.Information{
		OrderId:           orderId,
		Status:            order.StatusNotPay,
		LogisticsPackages: make([]*order.Logistics, 0, 0),
		GoodsList:         make([]*order.Goods, 0, 0),
		Discounts:         make([]*order.Discount, 0, 0),
		Address: &order.Address{
			OrderId: orderId,
		},
		Payments: make([]*order.Payment, 0, 0),
		ModificationLogs: []*order.StatusModificationLog{
			{
				OrderId:           orderId,
				DestinationStatus: order.StatusNotPay,
			},
		},
		Comment: &order.Comment{
			OrderId:   orderId,
			Commented: false,
		},
		Time: databases.Time{
			BasicTimeFields: databases.BasicTimeFields{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	for _, option := range options {
		option.Option(orderData)
	}

	return orderData
}

// 创建订单
func CreateOrder(data *order.Information, isVip bool, orderType order.Type) (result bool, err error) {

	data.IsVip = isVip
	data.OrderType = order.TypeShop
	if orderType > 0 {
		data.OrderType = orderType
	}

	err = operation.ComputeGoodsPrice(data, isVip)
	if err != nil {
		logrus.WithField("order", data).
			WithField("error", err).
			Info("compute goods price failed")

		return false, err
	}

	err = operation.ComputeDiscountPrice(data)
	if err != nil {
		logrus.WithField("order", data).
			WithField("error", err).
			Info("compute discount price failed")

		return false, err
	}

	err = operation.ComputeFreight(data)
	if err != nil {
		logrus.WithField("order", data).
			WithField("error", err).
			Info("compute freight price failed")

		return false, err
	}

	err = operation.ComputePayable(data)
	if err != nil {
		logrus.WithField("order", data).
			WithField("error", err).
			Info("compute payable failed")

		return false, err
	}

	if data.Point > 0 {
		// 扣积分
		userApi := client.GetUserService()
		if userApi == nil {
			logrus.Error("get user service client is nil")
			return false, nil
		}

		modifyUserPointReply, modifyUserError := userApi.ModifyUserPoint(context.Background(), &protos.ModifyUserPointRequest{
			UserId:  data.UserId,
			Remark:  "购买商品",
			OrderId: data.OrderId,
			Point:   data.Point,
			Type:    1,
		})

		if modifyUserError != nil {

			logrus.WithField("order", data).
				WithField("error", modifyUserError).
				Info("modify user point failed")
			return false, nil
		}

		if modifyUserPointReply.Err != nil {
			logrus.WithField("order", data).
				WithField("error", modifyUserPointReply.Err).
				Info("modify user point reply failed")

			return false, nil
		}

		logrus.Info("modify user point success")

		// 改状态
		data.Status = order.StatusPaid
	}

	err = saveData(data)

	if err != nil {
		logrus.WithField("order", data).
			WithField("error", err).
			Info("save data failed")

		return false, err
	}

	return true, nil
}

func saveData(data *order.Information) (err error) {

	err = database.GetDB(constants.DatabaseConfigKey).Create(data).Error
	if err != nil {

		logrus.WithField("order", data).
			WithField("error", err).
			Error("create order record")
		return
	}

	return
}

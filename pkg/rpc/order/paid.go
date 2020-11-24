package order

import (
	"time"

	"dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/model/rebate"
	rebateProtos "dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/rpc/protos"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	orderPayment "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/payment"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/parameters"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/responseerrors"
	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (svc Service) PaidOrder(ctx context.Context, req *pb.PaidOrderRequest) (resp *pb.PaidOrderReply, err error) {

	resp = &pb.PaidOrderReply{}

	logrus.WithField("requestData", req).WithField("context", ctx).Info("requested")

	requestData, replyError := svc.getPaidOrderRequestData(req)
	if replyError != nil {
		resp.Error = replyError.PBError
		logrus.WithField("error", replyError).Info("validate parameters error")
		return
	}

	logger := logrus.WithField("requestData", requestData).WithField("context", ctx)

	orderData, getOrderError := query.GetOrder(requestData.OrderId)
	if getOrderError != nil {

		resp.Error = responseerrors.NewError(responseerrors.ErrorOrderNotExist).WithError(getOrderError).PBError
		logger.WithField("error", getOrderError).Error("get order error")
		return
	}

	switch orderData.Status {
	case order.StatusNotPay:

		break

	// 取消的订单暂时不支持自动拉起
	case order.StatusCancel:

		resp.Error = responseerrors.NewError(responseerrors.ErrorOrderStatusCannotPay).PBError
		logger.WithField("error", resp.Error).Warn("order cannot paid")
		return

	// 后面如果开放分批支付的话，需要修改这里为根据订单当前情况进行判断
	case order.StatusPaid:
		resp.Error = responseerrors.NewError(responseerrors.ErrorOrderWasPaid).PBError
		logger.WithField("error", resp.Error).Warn("duplicate notify")
		return

	default:
		resp.Error = responseerrors.NewError(responseerrors.ErrorOrderStatusCannotPay).PBError
		logger.WithField("error", resp.Error).Warn("order cannot paid")
		return
	}

	if paidOrderError := orderPayment.Paid(orderData, svc.changeToOrderPayment(requestData.Payment)); paidOrderError != nil {

		resp.Error = responseerrors.NewError(responseerrors.ErrorInternalError).WithError(err).PBError
		logger.WithField("error", paidOrderError).Error("paid order error")
		return
	}

	resp.Success = true
	// 添加返利内容
	rebateApi := client.GetRebateService()
	if rebateApi == nil {
		logger.Error("Get rebate Service error")
		return
	}

	logger.Info("start rebate !!")
	logger.WithField("GoodsList", orderData.GoodsList).Info("goodsList")
	// 遍历数据 写入 返利表
	for _, goods := range orderData.GoodsList {
		rebateData := &rebateProtos.RebateOrder{
			OrderId:   req.OrderId,
			UserId:    orderData.UserId,
			PaidPrice: req.Payment.PaidPrice,
			PaidTime:  time.Now().Format("2006-01-02 15:04:05"),
			ItemType:  rebate.ItemTypeGoods,
			ShopId:    orderData.ShopId,
		}
		if !validation.IsEmpty(goods.SkuId) {
			rebateData.ItemId = goods.SkuId
		} else {
			logger.WithField("rebateData", rebateData).Error("ItemId is nil")
		}

		logger.WithField("data", rebateData).Info("rebateData")
		// 直接怼数据进去
		rebateReply, rebateReplyErr := rebateApi.CreateRebateOrder(ctx, &rebateProtos.CreateRebateOrderRequest{
			RebateOrder: rebateData,
		})

		if rebateReplyErr != nil {
			logger.WithError(rebateReplyErr).Error("rebateReply error")
			return
		}

		if rebateReply.Err != nil {
			logger.WithField("reply err", rebateReply.Err).Error("rebateReply.Err is not nil")
			return
		}

	}

	logger.Info("stop rebate !!")

	return
}

func (Service) getPaidOrderRequestData(req *pb.PaidOrderRequest) (requestData *parameters.PaidOrderRequest, responseError *responseerrors.Error) {

	requestData = &parameters.PaidOrderRequest{
		PaidOrderRequest: *req,
		Payment:          &parameters.Payment{Payment: *req.Payment},
	}

	if err := requestData.Validate(); err != nil {

		err, ok := err.(*responseerrors.Error)
		if ok {
			return nil, err
		}
		return nil, responseerrors.NewError(responseerrors.ErrorParameterError).WithError(err)
	}

	return requestData, nil
}

func (Service) changeToOrderPayment(requestData *parameters.Payment) (payment *order.Payment) {

	paidTime := time.Unix(int64(requestData.PaidTime), 0)
	return &order.Payment{
		TransactionId:  requestData.TransactionId,
		PaidPrice:      requestData.PaidPrice,
		PaymentChannel: requestData.PaymentChannel,
		PaymentMode:    requestData.PaymentMode,
		PaymentProduct: requestData.PaymentProduct,
		PaidTime:       &paidTime,
	}
}

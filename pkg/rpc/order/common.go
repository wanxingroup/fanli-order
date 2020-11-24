package order

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/cancel"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/responseerrors"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (svc Service) getOrder(logger *logrus.Entry, orderId uint64) (*protos.Order, *protos.Error) {

	orderData, err := query.GetOrder(orderId)
	if err != nil {

		logger.WithError(err).Error("get order error")
		return nil, responseerrors.CopyError(responseerrors.ErrorOrderNotExist)
	}

	return svc.convertOrder(orderData), nil
}

func (svc Service) convertOrder(orderData *order.Information) *protos.Order {

	var paidTime uint64
	if orderData.PaidTime != nil {
		paidTime = uint64(orderData.PaidTime.Unix())
	}
	responseOrder := &protos.Order{
		OrderId:         orderData.OrderId,
		ShopId:          orderData.ShopId,
		UserId:          orderData.UserId,
		Status:          svc.convertOrderStatusToProtobuf(orderData.Status),
		GoodsTotalPrice: orderData.GoodsTotalAmount,
		Payable:         orderData.Payable,
		Freight:         uint32(orderData.Freight),
		Point:           orderData.Point,
		IsVip:           orderData.IsVip,
		Remark:          orderData.Remark,
		LogisticsList:   make([]*protos.Logistics, 0, len(orderData.LogisticsPackages)),
		GoodsList:       make([]*protos.Goods, 0, len(orderData.GoodsList)),
		DiscountList:    make([]*protos.Discount, 0, len(orderData.Discounts)),
		Address: &protos.Address{
			Province:     orderData.Address.Province,
			City:         orderData.Address.City,
			District:     orderData.Address.District,
			Street:       orderData.Address.Street,
			Address:      orderData.Address.Address,
			ReceiverName: orderData.Address.ReceiverName,
			Tel:          orderData.Address.Tel,
		},
		PaymentList:               make([]*protos.Payment, 0, len(orderData.Payments)),
		StatusModificationLogList: make([]*protos.StatusModificationLog, 0, len(orderData.ModificationLogs)),
		CreateTime:                uint64(orderData.CreatedAt.Unix()),
		UpdateTime:                uint64(orderData.UpdatedAt.Unix()),
		AutoCancelTime:            uint64(0),
		PaidTime:                  paidTime,
		RefundStatus:              svc.ConvertRefundStatusToProtobuf(orderData.RefundStatus),
		OrderType:                 svc.ConvertOrderTypeToProtobuf(orderData.OrderType),
	}

	for _, goods := range orderData.GoodsList {

		responseOrder.GoodsList = append(responseOrder.GoodsList, &protos.Goods{
			SkuId:    goods.SkuId,
			SpuId:    goods.GoodsId,
			Count:    goods.Count,
			Price:    goods.Price,
			VipPrice: goods.VipPrice,
			Point:    goods.Point,
		})
	}

	for _, logistics := range orderData.LogisticsPackages {

		responseOrder.LogisticsList = append(responseOrder.LogisticsList, &protos.Logistics{
			PackageId:        logistics.PackageId,
			LogisticsCompany: logistics.LogisticsCompany,
			ExpressNumber:    logistics.ExpressNumber,
		})
	}

	for _, discount := range orderData.Discounts {

		responseOrder.DiscountList = append(responseOrder.DiscountList, &protos.Discount{
			Type:          uint32(discount.DiscountType),
			ObjectId:      discount.DiscountObjectId,
			DiscountPrice: discount.DiscountPrice,
		})
	}

	for _, modificationLog := range orderData.ModificationLogs {

		responseOrder.StatusModificationLogList = append(responseOrder.StatusModificationLogList, &protos.StatusModificationLog{
			DestinationStatus: svc.convertOrderStatusToProtobuf(modificationLog.DestinationStatus),
			Time:              uint64(modificationLog.CreatedAt.Unix()),
		})
	}

	if orderData.Status == order.StatusNotPay {

		responseOrder.AutoCancelTime = uint64(cancel.AutoCancelTime(orderData.CreatedAt).Unix())
	}
	return responseOrder
}

func (svc Service) convertToProtoError(logger *logrus.Entry, err error) *protos.Error {

	if err == nil {
		return nil
	}

	logger.WithError(err).Info("validate request data error")
	if validationError, ok := err.(validation.Error); ok {

		errCode, convertCodeError := strconv.Atoi(validationError.Code())
		if convertCodeError != nil {
			logger.WithError(convertCodeError).Error("convert validation error code error")
			errCode = int(responseerrors.ErrorParameterError.GetErrorCode())
		}

		return &protos.Error{
			ErrorCode:                uint32(errCode),
			ErrorMessageForDeveloper: validationError.Message(),
			ErrorMessageForUser:      responseerrors.ErrorParameterError.GetErrorMessageForUser(),
		}
	}

	return &protos.Error{
		ErrorCode:                responseerrors.ErrorParameterError.GetErrorCode(),
		ErrorMessageForDeveloper: err.Error(),
		ErrorMessageForUser:      responseerrors.ErrorParameterError.GetErrorMessageForUser(),
	}
}

func (svc Service) convertOrderStatusToProtobuf(status order.Status) protos.OrderStatus {

	switch status {
	case order.StatusNotPay:
		return protos.OrderStatus_NotPay
	case order.StatusPaid:
		return protos.OrderStatus_Paid
	case order.StatusDelivered:
		return protos.OrderStatus_Delivered
	case order.StatusReceived:
		return protos.OrderStatus_Received
	case order.StatusClosed:
		return protos.OrderStatus_Closed
	case order.StatusCompleted:
		return protos.OrderStatus_Completed
	case order.StatusCancel:
		return protos.OrderStatus_Cancel
	}

	return protos.OrderStatus_Unknown
}

func (svc Service) ConvertRefundStatusToProtobuf(refundStatus order.RefundStatus) protos.RefundStatus {

	switch refundStatus {
	case order.RefundStatusApplying:
		return protos.RefundStatus_RefundStatusApplying
	case order.RefundStatusCompleted:
		return protos.RefundStatus_RefundStatusCompleted
	}

	return protos.RefundStatus_RefundStatusNormal
}

func (svc Service) ConvertOrderTypeToProtobuf(orderType order.Type) protos.OrderType {

	switch orderType {
	case order.TypeShop:
		return protos.OrderType_OrderTypeShop
	case order.TypeCard:
		return protos.OrderType_OrderTypeCard
	}

	return protos.OrderType_OrderTypeShop
}

func (svc Service) ConvertOrderTypeFromProtobuf(refundStatus protos.OrderType) order.Type {

	switch refundStatus {
	case protos.OrderType_OrderTypeShop:
		return order.TypeShop
	case protos.OrderType_OrderTypeCard:
		return order.TypeCard
	}

	return order.TypeShop
}

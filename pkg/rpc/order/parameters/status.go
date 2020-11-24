package parameters

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

const OrderStatusUnknown = 0

func ConvertProtobufOrderStatusToModel(status protos.OrderStatus) order.Status {

	switch status {
	case protos.OrderStatus_NotPay:
		return order.StatusNotPay
	case protos.OrderStatus_Paid:
		return order.StatusPaid
	case protos.OrderStatus_Delivered:
		return order.StatusDelivered
	case protos.OrderStatus_Received:
		return order.StatusReceived
	case protos.OrderStatus_Closed:
		return order.StatusClosed
	case protos.OrderStatus_Completed:
		return order.StatusCompleted
	case protos.OrderStatus_Cancel:
		return order.StatusCancel
	}

	return order.Status(OrderStatusUnknown)
}

func ConvertProtobufRefundStatusToModel(refundStatus protos.RefundStatus) order.RefundStatus {
	switch refundStatus {
	case protos.RefundStatus_RefundStatusApplying:
		return order.RefundStatusApplying
	case protos.RefundStatus_RefundStatusCompleted:
		return order.RefundStatusCompleted
	case protos.RefundStatus_RefundStatusReject:
		return order.RefundStatusReject
	}

	return order.RefundStatusNormal
}

func ConvertProtobufRefundStatusListToModel(refundStatusList []protos.RefundStatus) []order.RefundStatus {
	list := make([]order.RefundStatus, 0, len(refundStatusList))
	for _, status := range refundStatusList {
		list = append(list, ConvertProtobufRefundStatusToModel(status))

	}

	return list
}

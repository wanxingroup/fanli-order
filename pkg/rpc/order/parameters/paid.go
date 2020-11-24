package parameters

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/validator"

	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

type PaidOrderRequest struct {
	pb.PaidOrderRequest
	Payment *Payment
}

func (request PaidOrderRequest) Validate() error {

	return validator.NewWrapper(
		validator.ValidateUint64PositiveInteger(request.OrderId, "orderId"),
		request.Payment.Validate,
	).Validate()
}

type Payment struct {
	pb.Payment
}

func (payment Payment) Validate() error {

	return validator.NewWrapper(
		validator.ValidateString(payment.TransactionId, "TransactionId", validator.ItemNotEmptyLimit, validator.ItemNoLimit),
		validator.ValidateUint64PositiveInteger(payment.PaidPrice, "paidPrice"),
		validator.ValidateString(payment.PaymentChannel, "paymentChannel", validator.ItemNotEmptyLimit, validator.ItemNoLimit),
		validator.ValidateString(payment.PaymentMode, "paymentMode", validator.ItemNotEmptyLimit, validator.ItemNoLimit),
		validator.ValidateString(payment.PaymentProduct, "paymentProduct", validator.ItemNotEmptyLimit, validator.ItemNoLimit),
		validator.ValidateUint64PositiveInteger(payment.PaidTime, "paidTime"),
	).Validate()
}

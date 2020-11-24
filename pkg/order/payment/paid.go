package payment

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func Paid(data *order.Information, paymentInformation *order.Payment) error {

	return newOperation().
		WithOrder(data).
		WithPayment(paymentInformation).
		Paid()
}

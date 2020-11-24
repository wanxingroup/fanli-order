package payment

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/status"
)

type paidOperation struct {
	logger  *logrus.Entry
	order   *order.Information
	payment *order.Payment
}

func newOperation() *paidOperation {
	return &paidOperation{logger: logrus.NewEntry(logrus.New())}
}

func (operation *paidOperation) WithOrder(data *order.Information) *paidOperation {

	operation.order = data
	operation.logger = operation.logger.WithField("order", data)
	return operation
}

func (operation *paidOperation) WithPayment(payment *order.Payment) *paidOperation {

	operation.payment = payment
	operation.logger = operation.logger.WithField("payment", payment)
	return operation
}

// 支付操作
func (operation *paidOperation) Paid() (err error) {

	if err = operation.checkWasUpdatedPaymentRecord(); err != nil {
		return err
	}

	err = status.ChangeStatus(operation.order, order.StatusPaid)
	if err != nil {
		operation.logger.
			WithField("action", "change status").
			Error(err)
		return
	}

	operation.order.PaidTime = operation.payment.PaidTime

	err = database.GetDB(constants.DatabaseConfigKey).
		Model(&order.Information{}).
		Where(&order.Information{
			OrderId: operation.order.OrderId,
		}).
		Update(&order.Information{PaidTime: operation.payment.PaidTime}).
		Error

	if err != nil {

		operation.logger.WithError(err).Error("update paid time error")
		return
	}

	if saveError := operation.savePayment(); saveError != nil {

		operation.logger.
			WithField("action", "create payment record").
			Error(saveError)
		// 由于订单记录支付信息只是作为一个信息冗余，所以如果记录失败，只是记录错误记录，不影响后面流程操作。
		// 这里后面应该做个通知机制
	}

	go operation.triggerVirtualGoodsDelivery()

	return
}

// 保存支付信息
func (operation *paidOperation) savePayment() (err error) {

	if operation.order.Payments == nil {
		operation.order.Payments = make([]*order.Payment, 0, 1)
	}

	for _, payment := range operation.order.Payments {

		if payment.TransactionId != operation.payment.TransactionId {

			continue
		}

		payment.PaidPrice = operation.payment.PaidPrice
		payment.PaymentChannel = operation.payment.PaymentChannel
		payment.PaymentMode = operation.payment.PaymentMode
		payment.PaymentProduct = operation.payment.PaymentProduct
		payment.PaidTime = operation.payment.PaidTime
		payment.Status = order.PaymentStatusSucceed

		err = database.GetDB(constants.DatabaseConfigKey).Model(payment).Update(payment).Error
		return
	}

	err = operation.createPaymentRecord(operation.payment)

	return
}

func (operation *paidOperation) createPaymentRecord(paymentData *order.Payment) (err error) {
	err = copier.Copy(paymentData, operation.payment)
	if err != nil {
		return errors.Wrap(err, "save payment")
	}

	paymentData.OrderId = operation.order.OrderId
	operation.order.Payments = append(operation.order.Payments, paymentData)

	err = database.GetDB(constants.DatabaseConfigKey).Create(paymentData).Error
	return
}

// 检查是否有重复的支付记录
func (operation *paidOperation) checkWasUpdatedPaymentRecord() (err error) {

	for _, payment := range operation.order.Payments {

		if payment.TransactionId != operation.payment.TransactionId {

			continue
		}

		if payment.Status == order.PaymentStatusPaying ||
			payment.Status == order.PaymentStatusCanceled {

			break
		}

		operation.logger.
			WithField("action", "call paid duplicated").
			Warn(err)
		return errorDuplicatedTransactionId
	}

	return
}

// 触发虚拟商品发货
func (operation *paidOperation) triggerVirtualGoodsDelivery() {

	// 后续虚拟商品发货可以在这里触发
}

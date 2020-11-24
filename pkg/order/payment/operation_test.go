package payment

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/status"
)

func TestPaid(t *testing.T) {

	shopId := uint64(1000)
	userId := uint64(100001)
	tests := []struct {
		inputOrder    *order.Information
		existOrder    *order.Information
		inputPayment  *order.Payment
		existsPayment *order.Payment
		want          *order.Information
		err           error
	}{
		// 正向测试
		{
			inputOrder: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
				Payments: []*order.Payment{
					{
						OrderId:        10000,
						TransactionId:  "10000004",
						PaidPrice:      100,
						PaymentChannel: "fuyou",
						PaymentMode:    "app",
						PaymentProduct: "weixin",
						Status:         order.PaymentStatusPaying,
					},
				},
			},
			existOrder: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
				Payments: []*order.Payment{
					{
						OrderId:        10000,
						TransactionId:  "10000004",
						PaidPrice:      100,
						PaymentChannel: "fuyou",
						PaymentMode:    "app",
						PaymentProduct: "weixin",
						Status:         order.PaymentStatusPaying,
					},
				},
			},
			inputPayment: &order.Payment{
				OrderId:        10000,
				TransactionId:  "10000004",
				PaidPrice:      100,
				PaymentChannel: "fuyou",
				PaymentMode:    "app",
				PaymentProduct: "weixin",
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusPaid,
				OrderType:        order.TypeShop,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
				Payments: []*order.Payment{
					{
						OrderId:        10000,
						TransactionId:  "10000004",
						PaidPrice:      100,
						PaymentChannel: "fuyou",
						PaymentMode:    "app",
						PaymentProduct: "weixin",
						Status:         order.PaymentStatusSucceed,
					},
				},
			},
		},
		// 重复支付测试
		{
			inputOrder: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
				Payments: []*order.Payment{
					{
						OrderId:        10000,
						TransactionId:  "10000004",
						PaidPrice:      100,
						PaymentChannel: "fuyou",
						PaymentMode:    "app",
						PaymentProduct: "weixin",
						Status:         order.PaymentStatusSucceed,
					},
				},
			},
			inputPayment: &order.Payment{
				OrderId:        10000,
				TransactionId:  "10000004",
				PaidPrice:      100,
				PaymentChannel: "fuyou",
				PaymentMode:    "app",
				PaymentProduct: "weixin",
			},
			err: errorDuplicatedTransactionId,
		},
		// 订单状态不符合支付测试
		{
			inputOrder: &order.Information{
				OrderId:          10001,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusPaid,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
			},
			existOrder: &order.Information{
				OrderId:          10001,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusPaid,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
			},
			inputPayment: &order.Payment{
				OrderId:        10001,
				TransactionId:  "10000005",
				PaidPrice:      100,
				PaymentChannel: "fuyou",
				PaymentMode:    "app",
				PaymentProduct: "weixin",
			},
			err: status.ErrorCannotChangeToSpecifyStatus,
		},
		// 支付信息存储失败测试
		{
			inputOrder: &order.Information{
				OrderId:          10002,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
			},
			existOrder: &order.Information{
				OrderId:          10002,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
			},
			inputPayment: &order.Payment{
				OrderId:        10002,
				TransactionId:  "10000006",
				PaidPrice:      100,
				PaymentChannel: "superlooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
				PaymentMode:    "app",
				PaymentProduct: "weixin",
			},
			want: &order.Information{
				OrderId:          10002,
				ShopId:           shopId,
				UserId:           userId,
				OrderType:        order.TypeShop,
				Status:           order.StatusPaid,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
				Payments:         []*order.Payment{},
			},
		},
	}

	for _, test := range tests {

		var err error
		if test.existOrder != nil {

			err = database.GetDB(constants.DatabaseConfigKey).Create(test.existOrder).Error
			assert.Nil(t, err, test.inputOrder)
		}

		if test.existsPayment != nil {
			err = database.GetDB(constants.DatabaseConfigKey).Create(test.existsPayment).Error
			assert.Nil(t, err)
		}

		err = Paid(test.inputOrder, test.inputPayment)
		assert.Equal(t, test.err, err)

		if test.want != nil {

			actualOrder := &order.Information{}
			err = database.GetDB(constants.DatabaseConfigKey).
				Model(actualOrder).
				Preload("Payments").
				Where(order.Information{OrderId: test.inputOrder.OrderId}).
				First(actualOrder).Error
			assert.Nil(t, err)

			assert.False(t, actualOrder.CreatedAt.IsZero())
			assert.False(t, actualOrder.UpdatedAt.IsZero())

			test.want.CreatedAt = actualOrder.CreatedAt
			test.want.UpdatedAt = actualOrder.UpdatedAt

			paymentsTime := map[string]databases.Time{}
			for _, payment := range actualOrder.Payments {

				paymentsTime[payment.TransactionId] = payment.Time
			}
			for _, payment := range test.want.Payments {
				payment.Time = paymentsTime[payment.TransactionId]
			}

			assert.EqualValues(t, test.want, actualOrder)
		}
	}
}

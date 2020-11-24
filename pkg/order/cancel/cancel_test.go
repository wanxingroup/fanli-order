package cancel

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestCancelOrder(t *testing.T) {
	orderId := uint64(44444)
	tests := []struct {
		inputOrderId     uint64
		inputInformation *order.Information
		wantInformation  *order.Information
	}{
		{
			inputOrderId: orderId,
			inputInformation: &order.Information{
				OrderId:    44444,
				MerchantId: 100,
				ShopId:     44444,
				UserId:     44444,
				Status:     order.StatusNotPay,
			},
			wantInformation: &order.Information{
				OrderId:    44444,
				MerchantId: 100,
				ShopId:     1000,
				UserId:     100000,
				Status:     order.StatusCancel,
			},
		},
	}

	for _, test := range tests {
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).Create(test.inputInformation)
		err := Cancel(test.inputOrderId)
		if err != nil {
			continue
		}
		orderData := &order.Information{}
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).
			Where(order.Information{OrderId: test.inputOrderId}).First(orderData)
		assert.EqualValues(t, test.wantInformation.Status, orderData.Status)
	}
}

func TestAutoCancelTime(t *testing.T) {
	testTime := time.Now()
	tests := []struct {
		inputOrderCreateTime time.Time
		wantOrderCreateTime  time.Time
		err                  error
	}{
		{
			inputOrderCreateTime: testTime,
			wantOrderCreateTime:  testTime.Add(constants.AutoCancelOrderSeconds * time.Second),
			err:                  nil,
		},
	}
	for _, test := range tests {
		resultTime := AutoCancelTime(test.inputOrderCreateTime)
		assert.EqualValues(t, test.wantOrderCreateTime, resultTime)
	}
}

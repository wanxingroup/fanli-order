package job

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestCancelOrdersJob(t *testing.T) {
	orderId := uint64(555555)
	tests := []struct {
		inputOrderId     uint64
		inputInformation *order.Information
	}{
		{
			inputOrderId: orderId,
			inputInformation: &order.Information{
				OrderId: 555555,
				ShopId:  555555,
				UserId:  555555,
				Status:  order.StatusNotPay,
			},
		},
	}

	for _, test := range tests {
		test.inputInformation.CreatedAt = test.inputInformation.CreatedAt.Add(-1000 * time.Second)
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).Create(test.inputInformation)
		CancelOrdersJob()
	}

}

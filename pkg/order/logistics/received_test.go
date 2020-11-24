package logistics

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestReceived(t *testing.T) {

	tests := []struct {
		inputOrder *order.Information
		want       *order.Information
		err        error
	}{
		// 正向测试
		{
			inputOrder: &order.Information{
				OrderId: 1000001,
				ShopId:  1000001,
				UserId:  1000001,
				Status:  order.StatusDelivered,
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusPaid,
					},
					{
						Model: gorm.Model{
							ID: 3,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusDelivered,
					},
				},
			},

			want: &order.Information{
				OrderId:   1000001,
				ShopId:    1000001,
				UserId:    1000001,
				Status:    order.StatusReceived,
				OrderType: order.TypeShop,
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusPaid,
					},
					{
						Model: gorm.Model{
							ID: 3,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusDelivered,
					},
					{
						Model: gorm.Model{
							ID: 4,
						},
						OrderId:           1000001,
						DestinationStatus: order.StatusReceived,
					},
				},
			},
		},
		// 反向测试，状态不符合
		{
			inputOrder: &order.Information{
				OrderId: 1000002,
				ShopId:  1000002,
				UserId:  1000002,
				Status:  order.StatusPaid,
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           1000002,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           1000002,
						DestinationStatus: order.StatusPaid,
					},
				},
			},
			err: errorStatusCannotReceived,
		},
	}

	for _, test := range tests {

		assert.Nil(t, database.GetDB(constants.DatabaseConfigKey).Create(test.inputOrder).Error)
		assert.Equal(t, test.err, Received(test.inputOrder))

		if test.err != nil {
			continue
		}

		orderData := &order.Information{}
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).
			Preload("ModificationLogs").
			Where(order.Information{OrderId: test.inputOrder.OrderId}).First(orderData)

		test.want.Time = orderData.Time

		logsModel := make(map[uint]gorm.Model)
		for _, log := range orderData.ModificationLogs {
			logsModel[log.ID] = log.Model
		}
		for _, log := range test.want.ModificationLogs {
			log.Model = logsModel[log.ID]
			log.OrderId = orderData.OrderId
		}

		assert.EqualValues(t, test.want, orderData)
	}
}

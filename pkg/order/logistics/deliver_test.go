package logistics

import (
	"fmt"
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestDeliver(t *testing.T) {

	tests := []struct {
		inputOrder     *order.Information
		inputLogistics *order.Logistics
		want           *order.Information
		err            error
	}{
		// 正向测试
		{
			inputOrder: &order.Information{
				OrderId:           10000,
				ShopId:            1000,
				UserId:            100000,
				Status:            order.StatusPaid,
				LogisticsPackages: []*order.Logistics{},
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Count:   1,
					},
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000002,
						Count:   2,
					},
				},
				Discounts: []*order.Discount{
					{
						OrderId:          10000,
						SerialId:         0,
						DiscountType:     1,
						DiscountPrice:    10,
						DiscountObjectId: "100-100001",
					},
				},
				Address: &order.Address{
					OrderId:      10000,
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
				Payments: []*order.Payment{},
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           10000,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           10000,
						DestinationStatus: order.StatusPaid,
					},
				},
				Comment: &order.Comment{
					OrderId:   10000,
					Commented: false,
				},
			},
			inputLogistics: &order.Logistics{
				LogisticsCompany: "顺丰速递",
				ExpressNumber:    "292-292-20291-13",
			},
			want: &order.Information{
				OrderId:   10000,
				ShopId:    1000,
				UserId:    100000,
				Status:    order.StatusDelivered,
				OrderType: order.TypeShop,
				LogisticsPackages: []*order.Logistics{
					{
						LogisticsCompany: "顺丰速递",
						ExpressNumber:    "292-292-20291-13",
					},
				},
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Count:   1,
					},
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000002,
						Count:   2,
					},
				},
				Discounts: []*order.Discount{
					{
						OrderId:          10000,
						SerialId:         0,
						DiscountType:     1,
						DiscountPrice:    10,
						DiscountObjectId: "100-100001",
					},
				},
				Address: &order.Address{
					OrderId:      10000,
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
				Payments: []*order.Payment{},
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           10000,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           10000,
						DestinationStatus: order.StatusPaid,
					},
					{
						Model: gorm.Model{
							ID: 3,
						},
						OrderId:           10000,
						DestinationStatus: order.StatusDelivered,
					},
				},
				Comment: &order.Comment{
					OrderId:   10000,
					Commented: false,
				},
			},
		},
		// 反向测试，状态不符合
		{
			inputOrder: &order.Information{
				OrderId:           10001,
				ShopId:            1000,
				UserId:            100000,
				Status:            order.StatusNotPay,
				LogisticsPackages: []*order.Logistics{},
				GoodsList: []*order.Goods{
					{
						OrderId: 10001,
						GoodsId: 20000,
						SkuId:   2000001,
						Count:   1,
					},
				},
				Discounts: []*order.Discount{},
				Address: &order.Address{
					OrderId:      10001,
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
				Payments: []*order.Payment{},
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           10001,
						DestinationStatus: order.StatusNotPay,
					},
				},
				Comment: &order.Comment{
					OrderId:   10001,
					Commented: false,
				},
			},
			inputLogistics: &order.Logistics{
				LogisticsCompany: "顺丰速递",
				ExpressNumber:    "292-292-20291-13",
			},
			err: errorStatusCannotDeliver,
		},
		// 反向测试，创建物流单号过长
		{
			inputOrder: &order.Information{
				OrderId:           10002,
				ShopId:            1000,
				UserId:            100000,
				Status:            order.StatusPaid,
				LogisticsPackages: []*order.Logistics{},
				GoodsList: []*order.Goods{
					{
						OrderId: 10002,
						GoodsId: 20000,
						SkuId:   2000001,
						Count:   1,
					},
				},
				Discounts: []*order.Discount{},
				Address: &order.Address{
					OrderId:      10002,
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
				Payments: []*order.Payment{},
				ModificationLogs: []*order.StatusModificationLog{
					{
						Model: gorm.Model{
							ID: 1,
						},
						OrderId:           10002,
						DestinationStatus: order.StatusNotPay,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						OrderId:           10002,
						DestinationStatus: order.StatusPaid,
					},
				},
				Comment: &order.Comment{
					OrderId:   10002,
					Commented: false,
				},
			},
			inputLogistics: &order.Logistics{
				LogisticsCompany: "顺丰速递",
				ExpressNumber:    "superlooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
			},
			err: &mysql.MySQLError{Number: 0x57e, Message: "Data too long for column 'expressNumber' at row 1"},
		},
	}

	for _, test := range tests {

		assert.Nil(t, database.GetDB(constants.DatabaseConfigKey).Create(test.inputOrder).Error)
		_, resultError := Deliver(test.inputOrder, test.inputLogistics)
		assert.Equal(t, test.err, resultError)

		if test.err != nil {
			continue
		}

		orderData := &order.Information{}
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).
			Preload("LogisticsPackages").
			Preload("GoodsList").
			Preload("Discounts").
			Preload("LogisticsPackages").
			Preload("Address").
			Preload("Payments").
			Preload("ModificationLogs").
			Preload("Comment").
			Where(order.Information{OrderId: test.inputOrder.OrderId}).First(orderData)

		test.want.Time = orderData.Time
		test.want.Address.Time = orderData.Address.Time
		test.want.Comment.Time = orderData.Comment.Time

		discountsTime := make(map[string]databases.Time)
		for _, discount := range orderData.Discounts {
			discountsTime[fmt.Sprintf("%d_%d", discount.OrderId, discount.SerialId)] = discount.Time
		}
		for _, discount := range test.want.Discounts {
			discount.Time = discountsTime[fmt.Sprintf("%d_%d", discount.OrderId, discount.SerialId)]
			discount.OrderId = orderData.OrderId
		}

		goodsTime := make(map[string]databases.Time)
		for _, goods := range orderData.GoodsList {
			goodsTime[fmt.Sprintf("%d_%d", goods.OrderId, goods.SkuId)] = goods.Time
		}
		for _, goods := range test.want.GoodsList {
			goods.Time = goodsTime[fmt.Sprintf("%d_%d", goods.OrderId, goods.SkuId)]
			goods.OrderId = orderData.OrderId
		}

		logsModel := make(map[uint]gorm.Model)
		for _, log := range orderData.ModificationLogs {
			logsModel[log.ID] = log.Model
		}
		for _, log := range test.want.ModificationLogs {
			log.Model = logsModel[log.ID]
			log.OrderId = orderData.OrderId
		}

		paymentTime := make(map[string]databases.Time)
		for _, payment := range orderData.Payments {
			paymentTime[payment.TransactionId] = payment.Time
		}
		for _, payment := range test.want.Payments {
			payment.Time = paymentTime[payment.TransactionId]
			payment.OrderId = orderData.OrderId
		}

		logisticsPackageIds := make(map[string]uint64)
		for _, logistics := range orderData.LogisticsPackages {
			logisticsPackageIds[fmt.Sprintf("%s_%s", logistics.LogisticsCompany, logistics.ExpressNumber)] = logistics.PackageId
		}
		for _, logistics := range test.want.LogisticsPackages {
			logistics.PackageId = logisticsPackageIds[fmt.Sprintf("%s_%s", logistics.LogisticsCompany, logistics.ExpressNumber)]
			logistics.OrderId = orderData.OrderId
		}

		logisticsTime := make(map[uint64]databases.Time)
		for _, logistics := range orderData.LogisticsPackages {
			logisticsTime[logistics.PackageId] = logistics.Time
		}
		for _, logistics := range test.want.LogisticsPackages {
			logistics.Time = logisticsTime[logistics.PackageId]
		}

		assert.EqualValues(t, test.want, orderData)
	}
}

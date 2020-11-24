package create

import (
	"fmt"
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestNewOrder(t *testing.T) {

	tests := []struct {
		input []Option
		want  *order.Information
	}{
		{
			input: []Option{
				NewOptionMerchant(100),
				NewOptionShop(10000),
				NewOptionUser(1000),
			},
			want: &order.Information{
				RefundStatus:      order.RefundStatusNormal,
				MerchantId:        100,
				ShopId:            10000,
				UserId:            1000,
				Status:            order.StatusNotPay,
				LogisticsPackages: make([]*order.Logistics, 0, 0),
				GoodsList:         make([]*order.Goods, 0, 0),
				Discounts:         make([]*order.Discount, 0, 0),
				Address:           &order.Address{},
				Payments:          make([]*order.Payment, 0, 0),
				ModificationLogs: []*order.StatusModificationLog{
					{
						DestinationStatus: order.StatusNotPay,
					},
				},
				Comment: &order.Comment{
					Commented: false,
				},
				Time: databases.Time{
					BasicTimeFields: databases.BasicTimeFields{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
		},
	}

	for _, test := range tests {

		orderData := NewOrder(0, test.input...)

		assert.Greater(t, orderData.OrderId, uint64(0))
		test.want.OrderId = orderData.OrderId

		for _, logistics := range test.want.LogisticsPackages {
			logistics.OrderId = orderData.OrderId
		}
		for _, goods := range test.want.GoodsList {
			goods.OrderId = orderData.OrderId
		}
		for _, discount := range test.want.Discounts {
			discount.OrderId = orderData.OrderId
		}
		for _, modificationLog := range test.want.ModificationLogs {
			modificationLog.OrderId = orderData.OrderId
		}
		test.want.Address.OrderId = orderData.OrderId
		test.want.Comment.OrderId = orderData.OrderId
		test.want.Time = orderData.Time
		assert.EqualValues(t, test.want, orderData)
	}
}

func TestCreateOrder(t *testing.T) {

	tests := []struct {
		input                *order.Information
		wantInformation      *order.Information
		wantAddress          *order.Address
		wantDiscounts        []*order.Discount
		wantGoodsList        []*order.Goods
		wantModificationLogs []*order.StatusModificationLog
		err                  error
	}{
		{
			input: &order.Information{
				OrderId:           10000,
				MerchantId:        100,
				ShopId:            1000,
				UserId:            100000,
				Status:            order.StatusNotPay,
				LogisticsPackages: []*order.Logistics{},
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   10,
						Count:   1,
					},
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000002,
						Price:   10,
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
				},
				Comment: &order.Comment{
					OrderId:   10000,
					Commented: false,
				},
			},
			wantInformation: &order.Information{
				OrderType:         order.TypeShop,
				RefundStatus:      order.RefundStatusNormal,
				OrderId:           10000,
				MerchantId:        100,
				ShopId:            1000,
				UserId:            100000,
				Status:            order.StatusNotPay,
				LogisticsPackages: []*order.Logistics{},
				GoodsTotalAmount:  30,
				DiscountPrice:     10,
				Payable:           20,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   10,
						Count:   1,
					},
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000002,
						Price:   10,
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
				},
				Comment: &order.Comment{
					OrderId:   10000,
					Commented: false,
				},
			},
			wantAddress: &order.Address{
				OrderId:      10000,
				Province:     "上海市",
				City:         "上海市",
				District:     "闵行区",
				Street:       "虹桥枢纽",
				Address:      "航中路1号",
				ReceiverName: "Lucky",
				Tel:          "13800138000",
			},
			wantDiscounts: []*order.Discount{
				{
					OrderId:          10000,
					SerialId:         0,
					DiscountType:     1,
					DiscountPrice:    10,
					DiscountObjectId: "100-100001",
				},
			},
			wantGoodsList: []*order.Goods{
				{
					OrderId: 10000,
					GoodsId: 20000,
					SkuId:   2000001,
					Price:   10,
					Count:   1,
				},
				{
					OrderId: 10000,
					GoodsId: 20000,
					SkuId:   2000002,
					Price:   10,
					Count:   2,
				},
			},
			wantModificationLogs: []*order.StatusModificationLog{
				{
					Model: gorm.Model{
						ID: 1,
					},
					OrderId:           10000,
					DestinationStatus: order.StatusNotPay,
				},
			},
		},
	}

	for _, test := range tests {

		_, err := CreateOrder(test.input, false, order.TypeShop)
		assert.Equal(t, test.err, err)

		if err != nil {
			continue
		}

		orderData := &order.Information{}
		address := &order.Address{}
		discounts := make([]*order.Discount, 0, 0)
		goodsList := make([]*order.Goods, 0, 0)
		logs := make([]*order.StatusModificationLog, 0, 0)
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
			Where(order.Information{OrderId: test.input.OrderId}).First(orderData)
		database.GetDB(constants.DatabaseConfigKey).Model(&order.Address{}).Where(order.Address{OrderId: test.input.OrderId}).First(address)
		database.GetDB(constants.DatabaseConfigKey).Model(&order.Discount{}).Where(order.Discount{OrderId: test.input.OrderId}).Find(&discounts)
		database.GetDB(constants.DatabaseConfigKey).Model(&order.Goods{}).Where(order.Goods{OrderId: test.input.OrderId}).Find(&goodsList)
		database.GetDB(constants.DatabaseConfigKey).Model(&order.StatusModificationLog{}).Where(order.StatusModificationLog{OrderId: test.input.OrderId}).Find(&logs)

		test.wantInformation.Time = orderData.Time
		test.wantInformation.Comment.Time = orderData.Comment.Time

		test.wantAddress.Time = address.Time
		test.wantInformation.Address.Time = address.Time

		discountsTime := make(map[string]databases.Time)
		for _, discount := range discounts {
			discountsTime[fmt.Sprintf("%d_%d", discount.OrderId, discount.SerialId)] = discount.Time
		}
		for _, discount := range test.wantDiscounts {
			discount.Time = discountsTime[fmt.Sprintf("%d_%d", discount.OrderId, discount.SerialId)]
		}
		for _, discount := range test.wantInformation.Discounts {
			discount.Time = discountsTime[fmt.Sprintf("%d_%d", discount.OrderId, discount.SerialId)]
		}

		goodsTime := make(map[string]databases.Time)
		for _, goods := range goodsList {
			goodsTime[fmt.Sprintf("%d_%d", goods.OrderId, goods.SkuId)] = goods.Time
		}
		for _, goods := range test.wantGoodsList {
			goods.Time = goodsTime[fmt.Sprintf("%d_%d", goods.OrderId, goods.SkuId)]
		}
		for _, goods := range test.wantInformation.GoodsList {
			goods.Time = goodsTime[fmt.Sprintf("%d_%d", goods.OrderId, goods.SkuId)]
		}

		logsModel := make(map[uint]gorm.Model)
		for _, log := range logs {
			logsModel[log.ID] = log.Model
		}
		for _, log := range test.wantModificationLogs {
			log.Model = logsModel[log.ID]
		}
		for _, log := range test.wantInformation.ModificationLogs {
			log.Model = logsModel[log.ID]
		}

		assert.EqualValues(t, test.wantInformation, orderData)
		assert.EqualValues(t, test.wantAddress, address)
		assert.EqualValues(t, test.wantDiscounts, discounts)
		assert.EqualValues(t, test.wantGoodsList, goodsList)
		assert.EqualValues(t, test.wantModificationLogs, logs)
	}
}

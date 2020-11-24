package operation

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestComputeGoodsPrice(t *testing.T) {

	shopId := uint64(1000)
	userId := uint64(100001)

	tests := []struct {
		input *order.Information
		want  *order.Information
		err   error
	}{
		// 正向测试，单个商品
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 0,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
		},
		// 正向测试多个商品
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 0,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   3,
					},
				},
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 300,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   3,
					},
				},
			},
		},
		// 正向测试多种商品多个商品
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 0,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   3,
					},
					{
						OrderId: 10000,
						GoodsId: 20001,
						SkuId:   2000101,
						Price:   50,
						Count:   1,
					},
				},
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 350,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   3,
					},
					{
						OrderId: 10000,
						GoodsId: 20001,
						SkuId:   2000101,
						Price:   50,
						Count:   1,
					},
				},
			},
		},
	}

	for _, test := range tests {

		err := ComputeGoodsPrice(test.input, false)
		assert.Equal(t, test.err, err)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestComputeFreight(t *testing.T) {

	hasRecordShopId := uint64(1000)
	noRecordShopId := uint64(1001)
	userId := uint64(100001)

	tests := []struct {
		input *order.Information
		exist *merchant.FreightSetting
		want  *order.Information
		err   error
	}{
		// 正向测试，金额不足包邮
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
			exist: &merchant.FreightSetting{
				ShopId:         hasRecordShopId,
				IsFreePostage:  false,
				ConditionPrice: 150,
				Freight:        5,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          5,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
		},
		// 正向测试，金额满足包邮
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
			exist: &merchant.FreightSetting{
				ShopId:         hasRecordShopId,
				IsFreePostage:  false,
				ConditionPrice: 50,
				Freight:        5,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   100,
						Count:   1,
					},
				},
			},
		},
		// 正向测试，商家包邮
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 1,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   1,
						Count:   1,
					},
				},
			},
			exist: &merchant.FreightSetting{
				ShopId:         hasRecordShopId,
				IsFreePostage:  true,
				ConditionPrice: 0,
				Freight:        0,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           hasRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 1,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   1,
						Count:   1,
					},
				},
			},
		},
		// 正向测试，无邮费配置记录商家
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           noRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 1,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   1,
						Count:   1,
					},
				},
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           noRecordShopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 1,
				Payable:          0,
				Freight:          0,
				GoodsList: []*order.Goods{
					{
						OrderId: 10000,
						GoodsId: 20000,
						SkuId:   2000001,
						Price:   1,
						Count:   1,
					},
				},
			},
		},
	}

	for _, test := range tests {

		if test.exist != nil {

			var count int
			assert.Nil(t, database.GetDB(constants.DatabaseConfigKey).
				Model(test.exist).
				Where(merchant.FreightSetting{ShopId: test.exist.ShopId}).
				Count(&count).
				Error,
			)
			if count == 0 {
				database.GetDB(constants.DatabaseConfigKey).Model(test.exist).Create(test.exist)
			} else {
				database.GetDB(constants.DatabaseConfigKey).Model(test.exist).Updates(test.exist)
			}
		}

		err := ComputeFreight(test.input)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.want, test.input)
	}
}

func TestComputePayable(t *testing.T) {

	shopId := uint64(1000)
	userId := uint64(100001)

	tests := []struct {
		input *order.Information
		want  *order.Information
		err   error
	}{
		// 正向测试，单个商品
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          100,
				Freight:          0,
			},
		},
		// 正向测试，有折扣优惠
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				DiscountPrice:    10,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          90,
				Freight:          0,
				DiscountPrice:    10,
			},
		},
		// 正向测试，有邮费
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 50,
				Payable:          0,
				Freight:          5,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 50,
				Payable:          55,
				Freight:          5,
			},
		},
		// 正向测试，有折扣优惠，有邮费
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 80,
				DiscountPrice:    10,
				Payable:          0,
				Freight:          5,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 80,
				DiscountPrice:    10,
				Payable:          75,
				Freight:          5,
			},
		},
		// 正向测试，有折扣优惠，有邮费，但是优惠更多，免费订单
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 5,
				DiscountPrice:    20,
				Payable:          0,
				Freight:          5,
			},
			want: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 5,
				DiscountPrice:    20,
				Payable:          0,
				Freight:          5,
			},
		},
	}

	for _, test := range tests {
		err := ComputePayable(test.input)
		assert.Equal(t, test.err, err, test)
		assert.EqualValues(t, test.want, test.input, test)
	}
}

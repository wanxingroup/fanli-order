package create

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestOptionShop_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input:       &order.Information{},
			inputOption: OptionShop{shopId: 10000},
			want: &order.Information{
				ShopId: 10000,
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionShop(t *testing.T) {

	tests := []struct {
		input uint64
		want  Option
	}{
		{
			input: 10000,
			want:  &OptionShop{shopId: 10000},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionShop(test.input))
	}
}

func TestOptionMerchant_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input:       &order.Information{},
			inputOption: OptionMerchant{merchantId: 100},
			want: &order.Information{
				MerchantId: 100,
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionMerchant(t *testing.T) {

	tests := []struct {
		input uint64
		want  Option
	}{
		{
			input: 100,
			want:  &OptionMerchant{merchantId: 100},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionMerchant(test.input))
	}
}

func TestOptionUser_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input:       &order.Information{},
			inputOption: OptionUser{userId: 10000},
			want: &order.Information{
				UserId: 10000,
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionUser(t *testing.T) {

	tests := []struct {
		input uint64
		want  Option
	}{
		{
			input: 10000,
			want:  &OptionUser{userId: 10000},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionUser(test.input))
	}
}

func TestOptionFreight_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input:       &order.Information{},
			inputOption: OptionFreight{freight: 10000},
			want: &order.Information{
				Freight: 10000,
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionFreight(t *testing.T) {

	tests := []struct {
		input uint64
		want  Option
	}{
		{
			input: 10000,
			want:  &OptionFreight{freight: 10000},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionFreight(test.input))
	}
}

func TestOptionGoods_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input: &order.Information{},
			inputOption: OptionGoods{
				goodsList: []*order.Goods{
					{
						GoodsId: 10000,
						SkuId:   1000001,
						Count:   1,
					},
				},
			},
			want: &order.Information{
				GoodsList: []*order.Goods{
					{
						GoodsId: 10000,
						SkuId:   1000001,
						Count:   1,
					},
				},
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionGoods(t *testing.T) {

	tests := []struct {
		input []*order.Goods
		want  Option
	}{
		{
			input: []*order.Goods{
				{
					GoodsId: 10000,
					SkuId:   1000001,
					Count:   1,
				},
			},
			want: &OptionGoods{
				goodsList: []*order.Goods{
					{
						GoodsId: 10000,
						SkuId:   1000001,
						Count:   1,
					},
				},
			},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionGoods(test.input...))
	}
}

func TestOptionDiscount_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input: &order.Information{},
			inputOption: OptionDiscount{
				discounts: []*order.Discount{
					{
						DiscountType:     1,
						DiscountObjectId: "100-10001",
					},
				},
			},
			want: &order.Information{
				Discounts: []*order.Discount{
					{
						DiscountType:     1,
						DiscountObjectId: "100-10001",
					},
				},
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionDiscount(t *testing.T) {

	tests := []struct {
		input []*order.Discount
		want  Option
	}{
		{
			input: []*order.Discount{
				{
					DiscountType:     1,
					DiscountObjectId: "100-10001",
				},
			},
			want: &OptionDiscount{
				discounts: []*order.Discount{
					{
						DiscountType:     1,
						DiscountObjectId: "100-10001",
					},
				},
			},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionDiscount(test.input...))
	}
}

func TestOptionAddress_Option(t *testing.T) {

	tests := []struct {
		input       *order.Information
		inputOption Option
		want        *order.Information
	}{
		{
			input: &order.Information{},
			inputOption: OptionAddress{
				address: &order.Address{
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
			},
			want: &order.Information{
				Address: &order.Address{
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
			},
		},
	}

	for _, test := range tests {

		test.inputOption.Option(test.input)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestNewOptionAddress(t *testing.T) {

	tests := []struct {
		input *order.Address
		want  Option
	}{
		{
			input: &order.Address{
				Province:     "上海市",
				City:         "上海市",
				District:     "闵行区",
				Street:       "虹桥枢纽",
				Address:      "航中路1号",
				ReceiverName: "Lucky",
				Tel:          "13800138000",
			},
			want: &OptionAddress{
				address: &order.Address{
					Province:     "上海市",
					City:         "上海市",
					District:     "闵行区",
					Street:       "虹桥枢纽",
					Address:      "航中路1号",
					ReceiverName: "Lucky",
					Tel:          "13800138000",
				},
			},
		},
	}

	for _, test := range tests {

		assert.EqualValues(t, test.want, NewOptionAddress(test.input))
	}
}

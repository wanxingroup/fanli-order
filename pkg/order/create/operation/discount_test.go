package operation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestDiscountElement_Deduct(t *testing.T) {

	shopId := uint64(1000)
	userId := uint64(100001)

	tests := []struct {
		input *order.Information
		want  *order.Information
		err   *ErrorAction
	}{
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				Discounts: []*order.Discount{
					{
						OrderId:       10000,
						DiscountPrice: 0,
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
				Discounts: []*order.Discount{
					{
						OrderId:       10000,
						DiscountPrice: 0, // TODO 实现后要可以获取到数值
					},
				},
			},
		},
	}

	for _, test := range tests {

		resource := &DiscountElement{}
		err := resource.Deduct(test.input)
		assert.Equal(t, test.err, err)
		assert.EqualValues(t, test.want, test.input)
	}
}

func TestDiscountElement_Rollback(t *testing.T) {

	shopId := uint64(1000)
	userId := uint64(100001)

	tests := []struct {
		input *order.Information
		want  *order.Information
		err   *ErrorAction
	}{
		{
			input: &order.Information{
				OrderId:          10000,
				ShopId:           shopId,
				UserId:           userId,
				Status:           order.StatusNotPay,
				GoodsTotalAmount: 100,
				Payable:          0,
				Freight:          0,
				Discounts: []*order.Discount{
					{
						OrderId:       10000,
						DiscountPrice: 10,
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
				Discounts: []*order.Discount{
					{
						OrderId:       10000,
						DiscountPrice: 10,
					},
				},
			},
		},
	}

	for _, test := range tests {

		resource := &DiscountElement{}
		err := resource.Rollback(test.input)
		assert.Equal(t, test.err, err)
		assert.EqualValues(t, test.want, test.input)
	}
}

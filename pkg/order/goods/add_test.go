package goods

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestAddGoods(t *testing.T) {
	tests := []struct {
		inputInformation *order.Information
		wantGoods        []*order.Goods
		err              error
	}{
		{
			inputInformation: &order.Information{
				OrderId: 66666,
				ShopId:  1000,
			},
			wantGoods: []*order.Goods{{
				GoodsId: 1000,
				SkuId:   1000,
			},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		AddGoods(test.inputInformation, test.wantGoods)
	}

}

package comment

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		inputInformation *order.Information
		wantComment      []*Comment
		err              error
	}{
		{
			inputInformation: &order.Information{
				OrderId:    55555,
				MerchantId: 100,
				ShopId:     1000,
			},
			wantComment: []*Comment{{
				GoodsId: 1000,
				SKUID:   1000,
			},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		Order(test.inputInformation, test.wantComment)
	}

}

package freightSetting

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func TestUpdateFreightSetting(t *testing.T) {
	tests := []struct {
		input *merchant.FreightSetting
		want  *merchant.FreightSetting
		err   error
	}{
		{
			input: &merchant.FreightSetting{
				ShopId:         33333,
				IsFreePostage:  true,
				ConditionPrice: 1000,
				Freight:        1000,
			},
			want: &merchant.FreightSetting{
				ShopId:         33333,
				IsFreePostage:  true,
				ConditionPrice: 20000,
				Freight:        2000,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		err := AddFreightSetting(test.input)
		if !assert.Equal(t, test.err, err) {
			return
		}
		test.input.ConditionPrice = 20000
		test.input.Freight = 2000
		err = UpdateFreightSetting(test.input)
		if !assert.Nil(t, err) {
			return
		}
		freightSetting, err := GetFreightSetting(test.input.ShopId)
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, test.want.Freight, freightSetting.Freight) {
			return
		}
	}
}

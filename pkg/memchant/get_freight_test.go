package freightSetting

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func TestGetFreightSetting(t *testing.T) {
	tests := []struct {
		input *merchant.FreightSetting
		want  *merchant.FreightSetting
		err   error
	}{
		{
			input: &merchant.FreightSetting{
				ShopId:         22222,
				IsFreePostage:  true,
				ConditionPrice: 22222,
				Freight:        22222,
			},
			want: &merchant.FreightSetting{
				ShopId:         22222,
				IsFreePostage:  true,
				ConditionPrice: 22222,
				Freight:        22222,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		err := AddFreightSetting(test.input)
		if !assert.Equal(t, test.err, err) {
			return
		}
		freightSetting, err := GetFreightSetting(test.input.ShopId)
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, test.want, freightSetting) {
			return
		}
	}
}

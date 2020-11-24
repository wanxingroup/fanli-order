package freightSetting

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func TestAddFreightSetting(t *testing.T) {
	tests := []struct {
		input *merchant.FreightSetting
		want  *merchant.FreightSetting
		err   error
	}{
		{
			input: &merchant.FreightSetting{
				ShopId:         11111,
				IsFreePostage:  true,
				ConditionPrice: 11111,
				Freight:        11111,
			},
			want: &merchant.FreightSetting{
				ShopId:         11111,
				IsFreePostage:  true,
				ConditionPrice: 11111,
				Freight:        11111,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		err := AddFreightSetting(test.input)
		if !assert.Equal(t, test.err, err) {
			return
		}
		freightSetting := new(merchant.FreightSetting)
		err = database.GetDB(constants.DatabaseConfigKey).Model(&freightSetting).Where("ShopId=?", test.want.ShopId).First(freightSetting).Error
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, test.want, freightSetting) {
			return
		}
	}
}

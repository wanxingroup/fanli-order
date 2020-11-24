package freightSetting

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func GetFreightSetting(shopId uint64) (*merchant.FreightSetting, error) {
	freightSetting := new(merchant.FreightSetting)
	err := database.GetDB(constants.DatabaseConfigKey).Where("shopId = ?", shopId).First(freightSetting).Error
	return freightSetting, err
}

package freightSetting

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func AddFreightSetting(data *merchant.FreightSetting) error {
	return database.GetDB(constants.DatabaseConfigKey).Create(data).Error
}

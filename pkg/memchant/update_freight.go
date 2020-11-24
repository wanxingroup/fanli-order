package freightSetting

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
)

func UpdateFreightSetting(data *merchant.FreightSetting) error {
	// 后续加wrap或者日志记录。
	return database.GetDB(constants.DatabaseConfigKey).Model(data).Updates(data).Error
}

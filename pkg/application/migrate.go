package application

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func migrateDatabases() {

	log.GetLogger().Info("starting to migrate database")

	migrateDatabaseAndLogError(merchant.FreightSetting{})
	migrateDatabaseAndLogError(order.Address{})
	migrateDatabaseAndLogError(order.Comment{})
	migrateDatabaseAndLogError(order.Discount{})
	migrateDatabaseAndLogError(order.Goods{})
	migrateDatabaseAndLogError(order.Information{})
	migrateDatabaseAndLogError(order.InformationCancelCron{})
	migrateDatabaseAndLogError(order.Logistics{})
	migrateDatabaseAndLogError(order.Payment{})
	migrateDatabaseAndLogError(order.StatusModificationLog{})

	log.GetLogger().Info("migrate database succeed")
}

func migrateDatabaseAndLogError(object interface{}) {

	if err := database.GetDB(constants.DatabaseConfigKey).AutoMigrate(object).Error; err != nil {

		log.GetLogger().WithField("object", fmt.Sprintf("%T", object)).
			WithError(err).Error("migrate failed")
	}
}

package schedule

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/cancel"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

// 自动取消未支付订单
// 功能描述： 搜索已经到达过期时间的，并且状态为未支付的订单
func AutoCancelOrder() {

	var orders []*order.Information
	var count int
	const pageSize = 100
	var page = 1
	for {
		err := database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).
			Where("`createdAt` < ? AND `status` = ?", cancel.BeforeTimeToCreateOrderNeedCancel(), order.StatusNotPay).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&orders).
			Count(&count).
			Error
		if err != nil {

			log.GetLogger().WithError(err).Error("find expire orders error")
			return
		}

		log.GetLogger().WithField("count", len(orders)).WithField("remainCount", count).Info("got expire orders")

		var processedCount = 0
		for _, orderData := range orders {

			log.GetLogger().WithField("order", orderData).Debug("iterating order")
			if cancelOrderError := cancel.Cancel(orderData.OrderId); cancelOrderError != nil {

				log.GetLogger().WithField("error", cancelOrderError).Error("automatic cancel order error")
				continue
			}

			processedCount++
			log.GetLogger().WithField("processedCount", processedCount).Debug("processed count + 1")
		}

		if count == 0 {

			log.GetLogger().Debug("count 0 break")
			break
		}

		// 防止一直都获取第一页订单，但是订单全部都处理失败，导致一直死循环的问题
		if processedCount == 0 {

			if count == 0 {
				log.GetLogger().Debug("processedCount 0 break")
				break
			} else { // 如果当前页有获取到订单，但是缺没有处理，可能是订单都异常了，所以尝试获取下一页。
				log.GetLogger().Debug("processedCount 0 next page")
				page++
			}
		}

		log.GetLogger().Debug("next round")
	}
}

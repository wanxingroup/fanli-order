package query

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func Orders(status order.Status, userId, shopId uint64, fromId uint64, limit uint) (orderList []*order.Information, err error) {

	db := database.GetDB(constants.DatabaseConfigKey).Model(&order.Information{})
	if fromId > 0 {
		db = db.Where("OrderId < ?", fromId)
	}
	db = db.Where(getCondition(status, userId, shopId))
	err = db.Order("OrderId desc, OrderId").Limit(limit).Preload("GoodsList").Find(&orderList).Error
	return orderList, err
}

func getCondition(status order.Status, userId, shopId uint64) *order.Information {
	var condition = order.Information{ShopId: shopId}
	if userId != 0 {
		condition.UserId = userId
	}
	if status != 0 {
		condition.Status = status
	}
	return &condition
}

func ShopOrders(query *QueryOrders) (orderList []*order.Information, count uint64, err error) {

	table := database.GetDB(constants.DatabaseConfigKey).
		Model(&order.Information{}).
		Where(getCondition(query.Status, query.UserId, query.ShopId))

	if len(query.UserIds) > 0 {
		table = table.Where("`userId` IN (?)", query.UserIds)
	}

	if query.MerchantId > 0 {
		table = table.Where("`merchantId` = ?", query.MerchantId)
	}

	if len(query.OrderIds) > 0 {
		table = table.Where("`orderId` IN (?)", query.OrderIds)
	}

	if len(query.RefundStatusList) > 0 {
		table = table.Where("`refundStatus` IN  (?)", query.RefundStatusList)
	}

	if query.OderId != 0 {
		table = table.Where(order.Information{OrderId: query.OderId})
	}

	if query.LastOrderId != 0 {
		table = table.Where("`orderId` < ?", query.LastOrderId)
	}

	if query.CreateTime != nil && !query.CreateTime.StartTime.IsZero() {
		table = table.Where("createdAt >= ?", query.CreateTime.StartTime)
	}

	if query.CreateTime != nil && !query.CreateTime.EndTime.IsZero() {
		table = table.Where("createdAt <= ?", query.CreateTime.EndTime)
	}

	if query.PaidTime != nil && !query.PaidTime.StartTime.IsZero() {
		table = table.Where("createdAt >= ?", query.CreateTime.StartTime)
	}

	if query.PaidTime != nil && !query.PaidTime.EndTime.IsZero() {
		table = table.Where("createdAt <= ?", query.CreateTime.EndTime)
	}

	if query.Page > 0 {
		table = table.Offset((query.Page - 1) * query.PageSize)
	}

	err = table.Order("orderId desc").Limit(query.PageSize).
		Preload("GoodsList").
		Preload("LogisticsPackages").
		Preload("Address").
		Find(&orderList).
		Error

	if err != nil {
		return
	}
	err = table.Count(&count).Error

	return orderList, count, err
}

func OrderStatusNums(userId, shopId uint64) ([]OrderStatus, error) {
	orderStatus := []OrderStatus{}
	var condition = order.Information{ShopId: shopId}
	if userId != 0 {
		condition = order.Information{UserId: userId, ShopId: shopId}
	}
	err := database.GetDB(constants.DatabaseConfigKey).
		Model(&order.Information{}).Where(condition).Group("status").
		Select("status as status,count(orderId) as orderNum").Scan(&orderStatus).Error
	return orderStatus, err
}

type OrderStatus struct {
	Status   string `gorm:"column:status" json:"status"`     // 订单状态
	OrderNum string `gorm:"column:orderNum" json:"orderNum"` // 订单数量
}

func GetOrder(orderId uint64) (orderData *order.Information, err error) {

	orderData = &order.Information{}
	err = database.GetDB(constants.DatabaseConfigKey).
		Model(&order.Information{}).
		Preload("LogisticsPackages").
		Preload("GoodsList").
		Preload("Discounts").
		Preload("Address").
		Preload("Payments").
		Preload("ModificationLogs").
		Preload("Comment").
		Where(order.Information{OrderId: orderId}).First(orderData).Error
	return
}

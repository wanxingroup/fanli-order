package operation

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/merchant"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func ComputeGoodsPrice(data *order.Information, isVip bool) error {

	// TODO: 现只支持单个商品购买，购物车商品结算需要重构
	goodsTotalAmount := uint64(0)
	for _, goods := range data.GoodsList {
		if goods.Point > 0 {
			data.Point += goods.Point
		} else {
			if isVip && goods.VipPrice > 0 {
				goodsTotalAmount += goods.VipPrice * uint64(goods.Count)
			} else {
				goodsTotalAmount += goods.Price * uint64(goods.Count)
			}
		}
	}

	data.GoodsTotalAmount = goodsTotalAmount
	return nil
}

func ComputeFreight(data *order.Information) error {

	freightSetting := &merchant.FreightSetting{}
	err := database.GetDB(constants.DatabaseConfigKey).Where(merchant.FreightSetting{ShopId: data.ShopId}).First(freightSetting).Error
	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return nil
		}

		return errors.Wrap(err, "get merchant freight error")
	}

	if freightSetting.IsFreePostage {
		return nil
	}

	if data.GoodsTotalAmount-data.DiscountPrice >= freightSetting.ConditionPrice {
		return nil
	}

	data.Freight = freightSetting.Freight
	return nil
}

func ComputePayable(data *order.Information) error {

	if data.GoodsTotalAmount+data.Freight < data.DiscountPrice {
		data.Payable = 0
	} else {
		data.Payable = data.GoodsTotalAmount - data.DiscountPrice + data.Freight
	}

	return nil
}

func ComputeDiscountPrice(data *order.Information) error {

	data.DiscountPrice = uint64(0)

	for _, discount := range data.Discounts {
		data.DiscountPrice += discount.DiscountPrice
	}

	return nil
}

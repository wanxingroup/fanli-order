package create

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Option interface {
	Option(order *order.Information)
}

type OptionShop struct {
	shopId uint64
}

func (option OptionShop) Option(order *order.Information) {

	order.ShopId = option.shopId
}

func NewOptionShop(shopId uint64) Option {

	return &OptionShop{shopId: shopId}
}

type OptionMerchant struct {
	merchantId uint64
}

func (option OptionMerchant) Option(order *order.Information) {

	order.MerchantId = option.merchantId
}

func NewOptionMerchant(merchantId uint64) Option {

	return &OptionMerchant{merchantId: merchantId}
}

type OptionUser struct {
	userId uint64
}

func (option OptionUser) Option(order *order.Information) {

	order.UserId = option.userId
}

func NewOptionUser(userId uint64) Option {

	return &OptionUser{userId: userId}
}

type OptionRemark struct {
	remark string
}

func (option OptionRemark) Option(order *order.Information) {

	order.Remark = option.remark
}

func NewOptionRemark(remark string) Option {

	return &OptionRemark{remark: remark}
}

type OptionFreight struct {
	freight uint64
}

func (option OptionFreight) Option(order *order.Information) {

	order.Freight = option.freight
}

func NewOptionFreight(freight uint64) Option {

	return &OptionFreight{freight: freight}
}

type OptionGoods struct {
	goodsList []*order.Goods
}

func (option OptionGoods) Option(order *order.Information) {

	order.GoodsList = option.goodsList
	for _, goods := range order.GoodsList {

		goods.OrderId = order.OrderId
	}
}

func NewOptionGoods(goodsList ...*order.Goods) Option {

	return &OptionGoods{goodsList: goodsList}
}

type OptionDiscount struct {
	discounts []*order.Discount
}

func (option OptionDiscount) Option(order *order.Information) {

	order.Discounts = option.discounts
	for _, discount := range order.Discounts {

		discount.OrderId = order.OrderId
	}
}

func NewOptionDiscount(discounts ...*order.Discount) Option {

	return &OptionDiscount{discounts: discounts}
}

type OptionAddress struct {
	address *order.Address
}

func (option OptionAddress) Option(order *order.Information) {

	order.Address = option.address
	order.Address.OrderId = order.OrderId
}

func NewOptionAddress(address *order.Address) Option {

	return &OptionAddress{address: address}
}

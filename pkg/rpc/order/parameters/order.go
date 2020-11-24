package parameters

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/validator"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

type Address struct {
	pb.Address
}

func (address Address) Validate() error {

	return validator.NewWrapper(
		validator.ValidateString(address.Province, "province", validator.ItemNotEmptyLimit, constants.ProvinceLengthLimit),
		validator.ValidateString(address.City, "city", validator.ItemNotEmptyLimit, constants.CityLengthLimit),
		validator.ValidateString(address.District, "district", validator.ItemNotEmptyLimit, constants.DistrictLengthLimit),
		validator.ValidateString(address.Street, "street", validator.ItemNotEmptyLimit, constants.StreetLengthLimit),
		validator.ValidateString(address.Address.Address, "address", validator.ItemNotEmptyLimit, constants.AddressLengthLimit),
		validator.ValidateString(address.ReceiverName, "receiverName", validator.ItemNotEmptyLimit, constants.ReceiverNameLengthLimit),
		validator.ValidateString(address.Tel, "tel", validator.ItemNotEmptyLimit, constants.TelLengthLimit),
	).Validate()
}

type OrderGoodsBase struct {
	pb.CreateOrderGoods
}

func (goods OrderGoodsBase) Validate() error {

	return validator.NewWrapper(
		validator.ValidateUint64Range(goods.SkuId, "skuId", uint64(validator.ItemNotEmptyLimit), validator.Uint64Maximum),
		validator.ValidateUint32Range(goods.Count, "count", uint32(validator.ItemNotEmptyLimit), validator.Uint32Maximum),
	).Validate()
}

type OrderDiscountBase struct {
	pb.CreateOrderDiscount
}

func (discount OrderDiscountBase) Validate() error {

	return validator.NewWrapper(
		validator.ValidateString(discount.ObjectId, "objectId", validator.ItemNotEmptyLimit, constants.DiscountObjectIdLengthLimit),
		validator.ValidateUint32Range(discount.Type, "type", uint32(validator.ItemNotEmptyLimit), validator.Uint32Maximum),
	).Validate()
}

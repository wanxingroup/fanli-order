package order

import (
	"fmt"
	"strconv"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/create"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/responseerrors"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func (svc Service) CreateOrder(ctx context.Context, req *protos.CreateOrderRequest) (*protos.CreateOrderReply, error) {

	var (
		err    error
		logger = log.GetLogger().WithContext(ctx)
	)

	err = validateCreateOrderData(req)
	if err != nil {

		return &protos.CreateOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	orderData := create.NewOrder(req.GetOrderId(),
		create.NewOptionMerchant(req.GetMerchantId()),
		create.NewOptionShop(req.GetShopId()),
		create.NewOptionUser(req.GetUserId()),
		create.NewOptionGoods(changeToOrderGoods(req.GetGoodsList())...),
		create.NewOptionAddress(changeToOrderAddress(req.GetAddress())),
		create.NewOptionDiscount(changeToOrderDiscounts(req.GetDiscounts())...),
		create.NewOptionRemark(req.GetRemark()),
	)

	createOrderResult, createOrderResultErr := create.CreateOrder(orderData, req.GetIsVip(), svc.ConvertOrderTypeFromProtobuf(req.GetOrderType()))

	if createOrderResultErr != nil {
		logger.WithError(createOrderResultErr).Error("create order error")

		return &protos.CreateOrderReply{
			Error: &protos.Error{
				ErrorCode:                responseerrors.ErrorCodeCreateOrderError,
				ErrorMessageForDeveloper: createOrderResultErr.Error(),
				ErrorMessageForUser:      responseerrors.ErrorMessageCreateOrderError,
			},
		}, nil
	}

	if createOrderResult != true {
		logger.Info("createOrderResult false")
		return &protos.CreateOrderReply{
			Error: &protos.Error{
				ErrorCode:           responseerrors.ErrorCodePointNotEnough,
				ErrorMessageForUser: responseerrors.ErrorMessagePointNotEnough,
			},
		}, nil
	}

	replyOrder, replyError := svc.getOrder(logger, orderData.OrderId)
	if replyError != nil {
		logger.WithField("error", err).Error("get order error")
		return &protos.CreateOrderReply{
			Error: &protos.Error{
				ErrorCode:           responseerrors.ErrorCodeGetOrderError,
				ErrorMessageForUser: responseerrors.ErrorMessageGetOrderError,
			},
		}, nil
	}

	return &protos.CreateOrderReply{
		Order: replyOrder,
	}, nil
}

func validateCreateOrderData(req *protos.CreateOrderRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.MerchantId, validation.Required.
			ErrorObject(validation.NewError(strconv.Itoa(responseerrors.ErrorCodeShopIdInvalid), responseerrors.ErrorMessageShopIdInvalid)),
		),
		validation.Field(&req.ShopId, validation.Required.
			ErrorObject(validation.NewError(strconv.Itoa(responseerrors.ErrorCodeShopIdInvalid), responseerrors.ErrorMessageShopIdInvalid)),
		),
		validation.Field(&req.UserId, validation.Required.
			ErrorObject(validation.NewError(strconv.Itoa(responseerrors.ErrorCodeUserIdInvalid), responseerrors.ErrorMessageUserIdInvalid)),
		),
		validation.Field(&req.GoodsList, validation.Required, validation.Length(1, 0).
			ErrorObject(validation.NewError(strconv.Itoa(responseerrors.ErrorCodeGoodsListEmpty), responseerrors.ErrorMessageGoodsListEmpty)),
			validation.By(validateCreateOrderGoodsList()),
		),
		validation.Field(&req.Discounts, validation.By(validateDiscount())),
		validation.Field(&req.Address, validation.By(validateAddress())),
	)
}

func validateAddress() validation.RuleFunc {

	return func(value interface{}) error {
		address, ok := value.(*protos.Address)
		if !ok || address == nil {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeAddressStructInvalid), responseerrors.ErrorMessageAddressStructInvalid)
		}

		return validation.ValidateStruct(address,
			validation.Field(&address.Province, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.ProvinceLengthLimit)),
			validation.Field(&address.City, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.CityLengthLimit)),
			validation.Field(&address.District, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.DistrictLengthLimit)),
			validation.Field(&address.Street, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.StreetLengthLimit)),
			validation.Field(&address.Address, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.AddressLengthLimit)),
			validation.Field(&address.ReceiverName, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.ReceiverNameLengthLimit)),
			validation.Field(&address.Tel, validation.Required, validation.RuneLength(validator.ItemNotEmptyLimit, constants.TelLengthLimit)),
		)
	}
}

func validateCreateOrderGoodsList() validation.RuleFunc {

	return func(value interface{}) error {
		goodsList, ok := value.([]*protos.CreateOrderGoods)

		if !ok || goodsList == nil {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeGoodsListStructInvalid), responseerrors.ErrorMessageGoodsListStructInvalid)
		}

		if len(goodsList) == 0 {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeGoodsListEmpty), responseerrors.ErrorMessageGoodsListEmpty)
		}

		for index, goods := range goodsList {

			err := validation.Validate(goods,
				validation.By(validateCreateOrderGoods()),
			)
			if err != nil {
				validationError, ok := err.(validation.ErrorObject)
				if ok {
					err = validationError.AddParam("index", index)
				} else {
					err = validation.NewError(strconv.Itoa(responseerrors.ErrorCodeGoodsDataError),
						fmt.Sprintf(responseerrors.ErrorMessageGoodsDataError, index, err.Error()),
					)
				}
				return err
			}
		}

		return nil
	}
}

func validateCreateOrderGoods() validation.RuleFunc {

	return func(value interface{}) error {

		goods, ok := value.(*protos.CreateOrderGoods)
		if !ok || goods == nil {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeGoodsDataError),
				responseerrors.ErrorMessageDiscountListStructInvalid)
		}

		return validation.ValidateStruct(goods,
			validation.Field(&goods.GoodsId, validation.Required),
			validation.Field(&goods.SkuId, validation.Required),
			validation.Field(&goods.Count, validation.Required),
		)
	}
}

func validateDiscount() validation.RuleFunc {

	return func(value interface{}) error {

		if value == nil {
			return nil
		}

		discountList, ok := value.([]*protos.CreateOrderDiscount)
		if !ok {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeDiscountListStructInvalid),
				responseerrors.ErrorMessageDiscountListStructInvalid)
		}

		if len(discountList) == 0 {
			return nil
		}

		for index, discount := range discountList {
			err := validation.Validate(discount,
				validation.By(validateCreateOrderDiscount()),
			)

			if err != nil {
				validationError, ok := err.(validation.ErrorObject)
				if ok {
					err = validationError.AddParam("index", index)
				} else {
					err = validation.NewError(strconv.Itoa(responseerrors.ErrorCodeDiscountDataError),
						fmt.Sprintf(responseerrors.ErrorMessageDiscountDataError, index, err.Error()),
					)
				}
				return err
			}
		}

		return nil
	}
}

func validateCreateOrderDiscount() validation.RuleFunc {

	return func(value interface{}) error {

		discount, ok := value.(*protos.CreateOrderDiscount)
		if !ok || discount == nil {
			return validation.NewError(strconv.Itoa(responseerrors.ErrorCodeDiscountDataError),
				responseerrors.ErrorMessageDiscountDataStructureError)
		}

		return validation.ValidateStruct(discount,
			validation.Field(&discount.Type, validation.Required),
			validation.Field(&discount.ObjectId, validation.Required),
		)
	}
}

func changeToOrderDiscounts(requestData []*protos.CreateOrderDiscount) []*order.Discount {

	var discounts = make([]*order.Discount, 0, len(requestData))
	var serialId uint8

	for _, discount := range requestData {

		discounts = append(discounts, &order.Discount{
			SerialId:         serialId,
			DiscountType:     order.DiscountType(discount.GetType()),
			DiscountObjectId: discount.GetObjectId(),
		})
		serialId++
	}

	return discounts
}

func changeToOrderAddress(address *protos.Address) *order.Address {

	return &order.Address{
		Province:     address.GetProvince(),
		City:         address.GetCity(),
		District:     address.GetDistrict(),
		Street:       address.GetStreet(),
		Address:      address.GetAddress(),
		ReceiverName: address.GetReceiverName(),
		Tel:          address.GetTel(),
	}
}

func changeToOrderGoods(requestData []*protos.CreateOrderGoods) []*order.Goods {

	var goodsList = make([]*order.Goods, 0, len(requestData))

	for _, goods := range requestData {

		goodsList = append(goodsList, &order.Goods{
			GoodsId:  goods.GetGoodsId(),
			SkuId:    goods.GetSkuId(),
			Price:    goods.GetPrice(),
			VipPrice: goods.GetVipPrice(),
			Point:    goods.GetPoint(),
			Count:    goods.Count,
		})
	}

	return goodsList
}

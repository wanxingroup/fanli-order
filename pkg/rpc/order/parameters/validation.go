package parameters

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

var ShopIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constants.ErrorCodeShopIdRequired, constants.ErrorMessageShopIdRequired)),
}

var OrderIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constants.ErrorCodeOrderIdRequired, constants.ErrorMessageOrderIdRequired)),
}

var DeliveryPackageRules = []validation.Rule{
	validation.By(func(value interface{}) error {
		deliveryPackage, ok := value.(*protos.DeliveryPackage)
		if !ok || deliveryPackage == nil {
			return validation.NewError(constants.ErrorCodeDeliveryPackageInvalid, constants.ErrorMessageDeliveryPackageInvalid)
		}

		return validateDeliveryPackage(deliveryPackage)
	}),
}

var LogisticsCompanyRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constants.ErrorCodeLogisticsCompanyRequired, constants.ErrorMessageLogisticsCompanyRequired)),
}

var ExpressNumberRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constants.ErrorCodeExpressNumberRequired, constants.ErrorMessageExpressNumberRequired)),
}

func validateDeliveryPackage(deliveryPackage *protos.DeliveryPackage) error {
	return validation.ValidateStruct(deliveryPackage,
		validation.Field(&deliveryPackage.LogisticsCompany, LogisticsCompanyRules...),
		validation.Field(&deliveryPackage.ExpressNumber, ExpressNumberRules...),
	)
}

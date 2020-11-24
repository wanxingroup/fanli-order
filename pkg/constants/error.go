package constants

const (
	ErrorCodeOrderIdRequired    = "403003" // 订单 ID 为必填
	ErrorMessageOrderIdRequired = "订单 ID 为必填"
)

const (
	ErrorCodeDeliveryPackageInvalid    = "403018" // 请求的发货包裹信息无效
	ErrorMessageDeliveryPackageInvalid = "请求的发货包裹信息无效"
)

const (
	ErrorCodeLogisticsCompanyRequired    = "403019" // 发货的物流公司名称为必填
	ErrorMessageLogisticsCompanyRequired = "发货的物流公司名称为必填"
)

const (
	ErrorCodeExpressNumberRequired    = "403020" // 发货的物流单号为必填
	ErrorMessageExpressNumberRequired = "发货的物流单号为必填"
)

const (
	ErrorCodeShopIdRequired    = "403021" // 店铺 ID 为必填
	ErrorMessageShopIdRequired = "店铺 ID 为必填"
)

const (
	ErrorCodeOrderStatusNoLongerNeedStatus               = 403022 // 订单状态不满足当前操作需求
	ErrorMessageOrderStatusNoLongerNeedStatusToUser      = "订单状态不满足当前操作需求"
	ErrorMessageOrderStatusNoLongerNeedStatusToDeveloper = "请重新刷新一下订单信息，当前状态已经不支持这个操作"
)

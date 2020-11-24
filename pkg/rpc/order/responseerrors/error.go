package responseerrors

import (
	"fmt"

	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

const errorCodeNotImplementYet = 1000
const errorCodeOrderNotExist = 200005
const errorCodeOrderWasPaid = 200006
const errorCodeOrderStatusCannotPay = 200007

const errorCodeParameterError = 100003
const errorCodeInternalError = 100004

const ErrorCodeGetOrderError = 503004    // 获取订单失败
const ErrorCodeCreateOrderError = 503005 // 创建订单失败
const ErrorCodeDeliverError = 503006     // 订单发货错误
const ErrorCodeReceivedError = 503007    // 确认收货错误
const ErrorCodeCancelOrderError = 503008 // 取消订单失败
const ErrorCodeQueryOrderError = 503009  // 查询订单失败
const ErrorCodeStatusTypeError = 503010  // 状态类型错误
const ErrorCodeLimitTypeError = 403006   // limit类型错误
const ErrorCodeFromPosTypeError = 403007 // fromPos类型错误
const ErrorCodeOrderIdTypeError = 403008 // 订单号类型错误
const ErrorCodeTimeTypeError = 403009    // 订单时间类型错误

const ErrorCodeShopIdInvalid = 403001             // 错误店铺ID
const ErrorCodeUserIdInvalid = 403002             // 错误用户ID
const ErrorCodeOrderIdInvalid = 403003            // 错误订单ID
const ErrorCodeNoticeMessage = 403004             // 商品库存不足
const ErrorCodeGoodsListEmpty = 403012            // 商品列表为空，请添加需要购买的商品
const ErrorCodeAddressStructInvalid = 403013      // 请求传入的地址信息结构无效
const ErrorCodeGoodsListStructInvalid = 403014    // 请求传入的商品数据结构无效
const ErrorCodeGoodsDataError = 403015            // 请求传入的商品数据中，# {index} 商品信息有误
const ErrorCodeDiscountListStructInvalid = 403016 // 请求传入的折扣信息数据结构无效
const ErrorCodeDiscountDataError = 403017         // 请求传入的折扣信息数据中，# {index} 商品信息有误

const ErrorMessageShopIdInvalid = "错误店铺ID"
const ErrorMessageUserIdInvalid = "错误用户ID"
const ErrorMessageOrderIdInvalid = "错误订单ID"
const ErrorMessageNoticeMessage = "商品库存不足，请稍安勿躁"
const ErrorMessageGoodsListEmpty = "商品列表为空，请添加需要购买的商品"
const ErrorMessageAddressStructInvalid = "请求传入的地址信息结构无效"
const ErrorMessageGoodsListStructInvalid = "请求传入的商品数据结构无效"
const ErrorMessageGoodsDataError = "请求传入的商品数据中，#%d 商品信息有误%s"
const ErrorMessageGoodsDataStructureError = "，商品结构无法解释"
const ErrorMessageDiscountListStructInvalid = "请求传入的折扣信息数据结构无效"
const ErrorMessageDiscountDataError = "请求传入的折扣数据中，#%d 折扣信息有误%s"
const ErrorMessageDiscountDataStructureError = "，折扣信息结构无法解释"
const ErrorMessageCreateOrderError = "创建订单失败"
const ErrorMessageGetOrderError = "获取订单失败"

const (
	ErrorCodePointNotEnough    = 403018
	ErrorMessagePointNotEnough = "积分余额不足"
)

type Error struct {
	PBError *pb.Error
}

func NewError(err pb.Error) *Error {
	return &Error{PBError: &err}
}

func (err Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.PBError.ErrorCode, err.PBError.ErrorMessageForDeveloper)
}

func (err Error) WithErrorCode(errorCode uint32) *Error {
	err.PBError.ErrorCode = errorCode
	return &err
}

func (err Error) WithError(standardError error) *Error {
	err.PBError.ErrorMessageForDeveloper = err.Error()
	return &err
}

func (err Error) WithErrorMessageForDeveloper(message string) *Error {
	err.PBError.ErrorMessageForDeveloper = message
	return &err
}

func (err Error) WithErrorMessageForUser(message string) *Error {
	err.PBError.ErrorMessageForUser = message
	return &err
}

var ErrorNotImplementYet = pb.Error{
	ErrorCode:                errorCodeNotImplementYet,
	ErrorMessageForDeveloper: "The API didn't implement yet",
	ErrorMessageForUser:      "功能尚未开放",
}

var ErrorGoodsListEmpty = pb.Error{
	ErrorCode:                ErrorCodeGoodsListEmpty,
	ErrorMessageForDeveloper: "Goods list empty",
	ErrorMessageForUser:      ErrorMessageGoodsListEmpty,
}

var ErrorParameterError = pb.Error{
	ErrorCode:                errorCodeParameterError,
	ErrorMessageForDeveloper: "Input parameters error",
	ErrorMessageForUser:      "输入参数错误，请重新校验格式后再请求",
}

var ErrorInternalError = pb.Error{
	ErrorCode:                errorCodeInternalError,
	ErrorMessageForDeveloper: "Internal call error",
	ErrorMessageForUser:      "内部调用错误，请稍后再试",
}

var ErrorOrderNotExist = pb.Error{
	ErrorCode:                errorCodeOrderNotExist,
	ErrorMessageForDeveloper: "order not exist",
	ErrorMessageForUser:      "订单不存在，请确认订单ID是否正确",
}

var ErrorOrderWasPaid = pb.Error{
	ErrorCode:                errorCodeOrderWasPaid,
	ErrorMessageForDeveloper: "order was paid",
	ErrorMessageForUser:      "订单已支付",
}

var ErrorOrderStatusCannotPay = pb.Error{
	ErrorCode:                errorCodeOrderStatusCannotPay,
	ErrorMessageForDeveloper: "order status cannot to pay",
	ErrorMessageForUser:      "订单不符合可支付状态",
}

func CopyError(err pb.Error) *pb.Error {

	return &err
}

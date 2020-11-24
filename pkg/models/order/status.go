package order

type Status uint8

const (
	StatusNotPay    Status = 1 // 未支付
	StatusPaid      Status = 2 // 已支付
	StatusDelivered Status = 3 // 已发货
	StatusReceived  Status = 4 // 已收货
	StatusClosed    Status = 5 // 订单关闭
	StatusCompleted Status = 6 // 订单完成
	StatusCancel    Status = 7 // 订单取消
)

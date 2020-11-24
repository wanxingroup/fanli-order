package order

type RefundStatus uint8

const (
	RefundStatusNormal    RefundStatus = 0 // 正常
	RefundStatusApplying  RefundStatus = 1 // 申请中
	RefundStatusCompleted RefundStatus = 2 // 退款完成
	RefundStatusReject    RefundStatus = 3 // 退款拒绝

)

package order

type PaymentStatus uint8

const (
	PaymentStatusPaying   = 1
	PaymentStatusSucceed  = 2
	PaymentStatusFailed   = 3
	PaymentStatusCanceled = 4
)

package operation

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type DiscountElement struct {
}

func (DiscountElement) Deduct(data *order.Information) *ErrorAction {

	// TODO Lucky 接入营销系统后需要扣减
	return nil
}

func (DiscountElement) Rollback(data *order.Information) *ErrorAction {

	// TODO Lucky 还原营销系统资源
	return nil
}

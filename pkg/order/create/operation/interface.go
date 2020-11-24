package operation

import (
	"sync"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Interface func(data *order.Information, wg *sync.WaitGroup, err chan<- ErrorAction)

type Resource interface {
	Deduct(data *order.Information) *ErrorAction
	Rollback(data *order.Information) *ErrorAction
}

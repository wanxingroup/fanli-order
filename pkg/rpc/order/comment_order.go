package order

import (
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (svc Service) CommentOrder(ctx context.Context, req *protos.CommentOrderRequest) (*protos.CommentOrderReply, error) {

	// TODO Lucky 评价订单实现
	return nil, nil
}

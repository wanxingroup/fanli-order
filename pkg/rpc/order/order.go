package order

import (
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"golang.org/x/net/context"

	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func (svc Service) OrderDetails(ctx context.Context, req *pb.OrderDetailsRequest) (resp *pb.OrderDetailsReply, err error) {

	resp = &pb.OrderDetailsReply{}
	resp.Order, resp.Error = svc.getOrder(rpcLog.WithRequestId(ctx, log.GetLogger()), req.OrderId)
	return
}

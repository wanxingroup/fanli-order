package state

import (
	"context"
	"errors"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (service *Service) Ping(ctx context.Context, req *protos.PingRequest) (reply *protos.PingReply, err error) {

	if req == nil {
		return nil, errors.New("input parameters error")
	}

	reply = &protos.PingReply{}
	reply.Message = req.Message
	return
}

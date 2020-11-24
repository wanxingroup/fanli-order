package order

import (
	"fmt"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/logistics"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/parameters"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func (svc Service) ReceivedOrder(ctx context.Context, req *protos.ReceivedOrderRequest) (*protos.ReceivedOrderReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateReceivedRequestData(req)
	if err != nil {
		return &protos.ReceivedOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	orderData, err := query.GetOrder(req.GetOrderId())
	if err != nil {
		logger.WithError(err).Error("get order error")
		return &protos.ReceivedOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	err = logistics.Received(orderData)
	if err != nil {
		logger.WithError(err).Error("deliver order error")
		return &protos.ReceivedOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	return &protos.ReceivedOrderReply{
		Success: true,
	}, nil
}

func validateReceivedRequestData(req *protos.ReceivedOrderRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.OrderId, parameters.OrderIdRules...),
	)
}

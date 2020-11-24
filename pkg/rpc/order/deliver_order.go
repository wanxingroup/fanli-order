package order

import (
	"fmt"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/logistics"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/parameters"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func (svc Service) DeliverOrder(ctx context.Context, req *protos.DeliverOrderRequest) (*protos.DeliverOrderReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateDeliverRequestData(req)
	if err != nil {
		return &protos.DeliverOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	orderData, err := query.GetOrder(req.GetOrderId())
	if err != nil {
		logger.WithError(err).Error("get order error")
		return &protos.DeliverOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	_, err = logistics.Deliver(orderData, &order.Logistics{
		LogisticsCompany: req.GetDeliveryPackage().GetLogisticsCompany(),
		ExpressNumber:    req.GetDeliveryPackage().GetExpressNumber(),
	})
	if err != nil {
		logger.WithError(err).Error("deliver order error")
		return &protos.DeliverOrderReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	return &protos.DeliverOrderReply{
		Success: true,
	}, nil
}

func validateDeliverRequestData(req *protos.DeliverOrderRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.OrderId, parameters.OrderIdRules...),
		validation.Field(&req.DeliveryPackage, parameters.DeliveryPackageRules...),
	)
}

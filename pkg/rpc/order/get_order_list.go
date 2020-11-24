package order

import (
	"time"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/parameters"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

func (svc Service) GetOrderList(ctx context.Context, req *protos.GetOrderListRequest) (*protos.GetOrderListReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	err := validateGetOrderListRequestData(req)

	logger.WithField("Req = ", req).Info("GetOrderList Request Data")

	logger.WithField("RefundStatusList", parameters.ConvertProtobufRefundStatusListToModel(req.GetRefundStatusList())).Info("GetOrderList RefundStatus")
	orders, count, err := query.ShopOrders(&query.QueryOrders{
		OrderType:        order.Type(req.GetOrderType()),
		Status:           parameters.ConvertProtobufOrderStatusToModel(req.GetStatus()),
		RefundStatusList: parameters.ConvertProtobufRefundStatusListToModel(req.GetRefundStatusList()),
		UserIds:          req.GetUserIds(),
		OrderIds:         req.GetOrderIds(),
		ShopId:           req.GetShopId(),
		MerchantId:       req.GetMerchantId(),
		LastOrderId:      req.GetLastOrderId(),
		PageSize:         req.GetPageSize(),
		Page:             req.GetPage(),
		CreateTime:       svc.convertProtobufTimeRangeToQuery(req.GetCreateTime()),
		PaidTime:         svc.convertProtobufTimeRangeToQuery(req.GetPaidTime()),
	})

	if err != nil {
		logger.WithError(err).Error("get order list error")
		return &protos.GetOrderListReply{
			Error: svc.convertToProtoError(logger, err),
		}, nil
	}

	return &protos.GetOrderListReply{
		OrderList: svc.convertOrderList(orders),
		Count:     count,
	}, nil
}

func (svc Service) convertOrderList(orders []*order.Information) []*protos.Order {

	if len(orders) == 0 {
		return []*protos.Order{}
	}

	result := make([]*protos.Order, 0, len(orders))
	for _, orderData := range orders {

		result = append(result, svc.convertOrder(orderData))
	}

	return result
}

func (svc Service) convertProtobufTimeRangeToQuery(timeRange *protos.TimeRange) *query.TimeRange {

	if timeRange == nil {
		return nil
	}

	if timeRange.GetStartTime() == 0 && timeRange.GetEndTime() == 0 {
		return nil
	}

	result := &query.TimeRange{}
	if timeRange.GetStartTime() != 0 {
		result.StartTime = time.Unix(int64(timeRange.GetStartTime()), 0)
	}
	if timeRange.GetEndTime() != 0 {
		result.EndTime = time.Unix(int64(timeRange.GetEndTime()), 0)
	}

	return result
}

func validateGetOrderListRequestData(req *protos.GetOrderListRequest) error {

	//if req.PageSize > constants.MaxOrderListPageSize {
	//	req.PageSize = constants.MaxOrderListPageSize
	//} else if req.PageSize == 0 {
	//	req.PageSize = constants.DefaultOrderListPageSize
	//}

	return nil
}

package status

import (
	"fmt"

	fuyouConstant "dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway/pkg/constant"
	fuyouProtos "dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway/pkg/rpc/protos"
	goodsProtos "dev-gitlab.wanxingrowth.com/fanli/goods/v2/pkg/rpc/protos"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

type Cancel struct {
}

func (_ Cancel) CanChangeTo(status order.Status) bool {
	// Cancel是最终值
	return false
}

func (_ Cancel) GetStatus() order.Status {
	return order.StatusCancel
}

func (_ Cancel) OnPrevChange(order *order.Information) error {
	return nil
}

func (status Cancel) OnChanged(order *order.Information) error {

	if err := status.restoreGoodsStocks(order); err != nil {
		return err
	}

	if err := status.cancelOrder(order); err != nil {
		return err
	}

	return nil
}

func (_ Cancel) restoreGoodsStocks(orderData *order.Information) error {

	logger := log.GetLogger().WithField("order", orderData)

	// 还原库存  还原折扣
	if len(orderData.GoodsList) == 0 {
		logger.Debug("goods list is empty")
		return nil
	}

	service := client.GetSPUService()
	if service == nil {

		logger.Error("get spu service client is nil")
		return fmt.Errorf("get spu service client is nil")
	}

	for _, goods := range orderData.GoodsList {
		logger.WithFields(logrus.Fields{
			"skuId": goods.SkuId,
			"count": goods.Count,
		}).Debug("starting restore goods stock")

		reply, err := service.RestoreSkuStock(context.Background(), &goodsProtos.RestoreSkuStockRequest{
			SkuId: goods.SkuId,
			Count: uint64(goods.Count),
		})
		if err != nil {

			logger.WithError(err).WithField("goods", goods).Error("restore goods stock error")
			return err
		}

		if reply.Err != nil {

			logger.WithError(err).WithField("goods", goods).Error("restore goods stock reply error")
			return fmt.Errorf("errorCode: %d, errorMessage: %s", reply.Err.ErrorCode, reply.Err.ErrorMessageForUser)
		}
	}
	return nil
}

func (_ Cancel) cancelOrder(orderData *order.Information) error {

	logger := log.GetLogger().WithField("order", orderData)

	api := client.GetFuYouService()
	if api == nil {
		logger.Error("get payment gateway rpc client error")
		return fmt.Errorf("get payment gateway rpc client error")
	}

	reply, err := api.CloseTransaction(context.Background(), &fuyouProtos.CloseTransactionRequest{
		SourceId:   orderData.OrderId,
		SourceType: fuyouProtos.SourceType_OrderService,
	})
	if err != nil {

		logger.WithError(err).Error("cancel payment transaction error")
		return err
	}

	if reply.GetErr() != nil {

		// 由于远方没有运行过去支付接口，会出现没有交易流水的记录
		// 所以富友支付网关报错，导致订单取消失败
		// 因此如果碰到富友没有找到支付流水记录，就处理为正常即可
		if reply.GetErr().GetCode() == fuyouConstant.ErrorCodeTransactionNotExist {
			return nil
		}

		logger.WithField("replyError", reply.GetErr()).Error("cancel payment transaction reply error")
		return fmt.Errorf("errorCode: %d, errorMessage: %s", reply.GetErr().GetCode(), reply.GetErr().GetMessage())
	}

	return nil
}

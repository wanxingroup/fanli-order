package order

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/cancel"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/responseerrors"
	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (Service) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (resp *pb.CancelOrderReply, err error) {

	resp = &pb.CancelOrderReply{}
	if cancelOrderError := cancel.Cancel(req.OrderId); cancelOrderError != nil {

		if cancel.IsOrderNotExistError(cancelOrderError) {

			logrus.WithError(cancelOrderError).Info("order not exist")
			resp.Error = responseerrors.NewError(responseerrors.ErrorOrderNotExist).WithError(cancelOrderError).PBError
			return
		}

		if cancel.IsOrderStatusCannotCancel(cancelOrderError) {

			logrus.WithError(cancelOrderError).Info("order status can not cancel")
			resp.Error = &pb.Error{
				ErrorCode:                constants.ErrorCodeOrderStatusNoLongerNeedStatus,
				ErrorMessageForDeveloper: constants.ErrorMessageOrderStatusNoLongerNeedStatusToDeveloper,
				ErrorMessageForUser:      constants.ErrorMessageOrderStatusNoLongerNeedStatusToUser,
			}
			return
		}

		logrus.WithError(cancelOrderError).Info("cancel order error")
		resp.Error = responseerrors.NewError(responseerrors.ErrorInternalError).WithError(cancelOrderError).PBError
		return
	}

	resp.Success = true
	return
}

package order

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/order/query"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order/responseerrors"
	pb "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
)

func (Service) ChangeRefundStatus(ctx context.Context, req *pb.ChangeRefundStatusRequest) (resp *pb.ChangeRefundStatusReply, err error) {

	resp = &pb.ChangeRefundStatusReply{}
	information, err := query.GetOrder(req.OrderId)
	if err != nil {
		resp.Err = responseerrors.NewError(responseerrors.ErrorInternalError).WithError(err).PBError

		return
	}

	if information == nil {
		resp.Err = responseerrors.NewError(responseerrors.ErrorOrderNotExist).PBError
		return
	}

	if uint8(req.Status) < uint8(information.RefundStatus) {

		resp.Err = &pb.Error{
			ErrorCode:                constants.ErrorCodeOrderStatusNoLongerNeedStatus,
			ErrorMessageForDeveloper: constants.ErrorMessageOrderStatusNoLongerNeedStatusToDeveloper,
			ErrorMessageForUser:      constants.ErrorMessageOrderStatusNoLongerNeedStatusToUser,
		}
		return

	}

	err = setRefundStatus(information, order.RefundStatus(req.Status))
	if err != nil {
		resp.Err = responseerrors.NewError(responseerrors.ErrorInternalError).WithError(err).PBError
		return
	}

	resp.Success = true
	return
}

func setRefundStatus(data *order.Information, refundStatus order.RefundStatus) error {

	err := database.GetDB(constants.DatabaseConfigKey).
		Model(data).
		Where(order.Information{RefundStatus: data.RefundStatus}).
		Select("refundStatus").
		Updates(order.Information{RefundStatus: refundStatus}).
		Error

	if err != nil {

		return err
	}

	data.RefundStatus = refundStatus

	return nil
}

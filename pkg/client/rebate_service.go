package client

import (
	"context"

	rebateProtos "dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

var rebateRPCService rebateProtos.RebateControllerClient

func InitRebateService() {
	rpcName := constants.RPCRebateServiceConfigKey
	log.GetLogger().Info("starting init rebate rpc service")

	var ctx = context.Background()
	var rpcConfig, exist = config.Config.RPCServices[rpcName]
	if !exist {
		log.GetLogger().Error(rpcName + " rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error(rpcName + " rpc service connect failed")
		return
	}

	rebateRPCService = rebateProtos.NewRebateControllerClient(conn)

	log.GetLogger().Info(rpcName + " rpc service init succeed")
}

func GetRebateService() rebateProtos.RebateControllerClient {
	return rebateRPCService
}

func SetRebateMockClient(client rebateProtos.RebateControllerClient) {
	rebateRPCService = client
}

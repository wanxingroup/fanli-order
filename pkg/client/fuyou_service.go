package client

import (
	"context"

	"dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

var fuyouRPCService protos.PayControllerClient

func InitFuYouService() {

	log.GetLogger().Info("starting init fuyou rpc service")

	var ctx = context.Background()

	var rpcConfig, exist = config.Config.RPCServices[constants.RPCFuYouServiceConfigKey]
	if !exist {
		log.GetLogger().Error("fuyou rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error("fuyou rpc service connect failed")
		return
	}

	fuyouRPCService = protos.NewPayControllerClient(conn)

	log.GetLogger().Info("fuyou rpc service init succeed")
}

func GetFuYouService() protos.PayControllerClient {

	return fuyouRPCService
}

func SetFuYouMockClient(client protos.PayControllerClient) {
	fuyouRPCService = client
}

package client

import (
	"context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

var userRPCService protos.UserControllerClient

func InitUserService() {

	log.GetLogger().Info("starting init user rpc service")

	var ctx = context.Background()

	var rpcConfig, exist = config.Config.RPCServices[constants.RPCUserServiceConfigKey]
	if !exist {
		log.GetLogger().Error("user rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error("user rpc service connect failed")
		return
	}

	userRPCService = protos.NewUserControllerClient(conn)

	log.GetLogger().Info("user rpc service init succeed")
}

func GetUserService() protos.UserControllerClient {

	return userRPCService
}

func SetUserMockClient(client protos.UserControllerClient) {
	userRPCService = client
}

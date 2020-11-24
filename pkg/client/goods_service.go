package client

import (
	"context"

	goodsProtos "dev-gitlab.wanxingrowth.com/fanli/goods/v2/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

var spuRPCService goodsProtos.SPUClient
var categoryRPCService goodsProtos.CategoryClient

func InitGoodsService() {

	log.GetLogger().Info("starting init goods rpc service")

	var ctx = context.Background()
	var rpcConfig, exist = config.Config.RPCServices[constants.RPCGoodsServiceConfigKey]
	if !exist {
		log.GetLogger().Error("goods rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error("goods rpc service connect failed")
		return
	}

	spuRPCService = goodsProtos.NewSPUClient(conn)
	categoryRPCService = goodsProtos.NewCategoryClient(conn)

	log.GetLogger().Info("goods rpc service init succeed")
}

func GetSPUService() goodsProtos.SPUClient {

	return spuRPCService
}

func GetCategoryService() goodsProtos.CategoryClient {
	return categoryRPCService
}

func SetSPUMockClient(client goodsProtos.SPUClient) {

	spuRPCService = client
}

func SetCategoryService(client goodsProtos.CategoryClient) {

	categoryRPCService = client
}

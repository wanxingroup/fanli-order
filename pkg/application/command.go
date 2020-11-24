package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/launcher"
	idCreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/order"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/state"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/schedule"
	cronUtils "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/cron"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/log"
)

var c *cron.Cron

func Start() {

	app := launcher.NewApplication(
		launcher.SetApplicationDescription(
			&launcher.ApplicationDescription{
				ShortDescription: "order service",
				LongDescription:  "order logical service.",
			},
		),
		launcher.SetApplicationLogger(log.GetLogger()),
		launcher.SetApplicationEvents(
			launcher.NewApplicationEvents(
				launcher.SetOnInitEvent(func(app *launcher.Application) {

					unmarshalConfiguration()

					registerOrderRPCRouter(app)

					idCreator.InitCreator(app.GetServiceId())

					client.InitFuYouService()
					client.InitGoodsService()
					client.InitUserService()
					client.InitRebateService()

					registerCronJobInit()
				}),
				launcher.SetOnStartEvent(func(app *launcher.Application) {

					migrateDatabases()
					c.Start()
				}),
				launcher.SetOnCloseEvent(func(app *launcher.Application) {

					ctx := c.Stop()
					ctx.Done()
				}),
			),
		),
	)

	app.Launch()
}

func registerOrderRPCRouter(app *launcher.Application) {

	rpcService := app.GetRPCService()
	if rpcService == nil {

		log.GetLogger().WithField("stage", "onInit").Error("get rpc service is nil")
		return
	}

	protos.RegisterOrderControllerServer(rpcService.GetRPCConnection(), &order.Service{})
	protos.RegisterStatusControllerServer(rpcService.GetRPCConnection(), &state.Service{})
}

func unmarshalConfiguration() {
	err := viper.Unmarshal(config.Config)
	if err != nil {

		log.GetLogger().WithError(err).Error("unmarshal config error")
	}
}

func registerCronJobInit() {

	c = cron.New(
		cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)),
		cron.WithLogger(cronUtils.NewLogger(log.GetLogger())),
	)

	type Cron struct {
		Spec string
		Run  func()
	}

	jobs := []Cron{
		{Spec: "* * * * *", Run: schedule.AutoCancelOrder},
	}

	for _, job := range jobs {
		_, err := c.AddFunc(job.Spec, job.Run)
		if err != nil {
			log.GetLogger().WithError(err).Error("register cron job error")
		}
	}
}

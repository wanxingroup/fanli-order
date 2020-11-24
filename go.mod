module dev-gitlab.wanxingrowth.com/fanli/order/v2

go 1.13

require (
	dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway v0.0.0-20200825020556-9f69154208da
	dev-gitlab.wanxingrowth.com/fanli/goods/v2 v2.0.7
	dev-gitlab.wanxingrowth.com/fanli/rebate v0.0.0-20200928130840-0530b217daf7
	dev-gitlab.wanxingrowth.com/fanli/user v0.0.2
	dev-gitlab.wanxingrowth.com/wanxin-go-micro/base v0.2.26
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/coreos/go-semver v0.3.0
	github.com/go-ozzo/ozzo-validation/v4 v4.2.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jinzhu/gorm v1.9.12
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/swag v1.6.7
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/grpc v1.30.0
)

replace dev-gitlab.wanxingrowth.com/wanxin-go-micro/base => github.com/wanxingroup/base v0.2.27

replace dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway => github.com/wanxingroup/fanli-fuyou-payment-gateway v0.0.0-20200828

replace dev-gitlab.wanxingrowth.com/fanli/goods/v2 => github.com/wanxingroup/fanli-goods/v2 v2.0.0-20201124070303-ea0a037380c1

replace dev-gitlab.wanxingrowth.com/fanli/rebate => github.com/wanxingroup/fanli-rebate v0.0.0

replace dev-gitlab.wanxingrowth.com/fanli/user => github.com/wanxingroup/fanli-user v0.0.2

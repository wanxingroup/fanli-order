package comment

import (
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

type Comment struct {
	GoodsId uint64 // 商品ID
	SKUID   uint64 // SKU ID
	Point   uint8  // 评分
	Content string // 评价内容
}

func Order(data *order.Information, comments []*Comment) error {

	// TODO Lucky 评论订单处理
	return nil
}

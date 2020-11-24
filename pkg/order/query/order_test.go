package query

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/constants"
	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/models/order"
)

func TestOrders(t *testing.T) {
	tests := []struct {
		inputStatus order.Status
		inputUserId uint64
		inputShopId uint64
		inputFromId uint64
		inputLimit  uint
		err         error
	}{
		{
			inputStatus: order.StatusNotPay,
			inputUserId: 111,
			inputShopId: 111,
			inputFromId: 111,
			inputLimit:  1,
			err:         nil,
		},
	}
	for _, test := range tests {
		_, err := Orders(test.inputStatus, test.inputUserId, test.inputShopId, test.inputFromId, test.inputLimit)
		if !assert.Equal(t, test.err, err) {
			return
		}
	}
}

func TestGetCondition(t *testing.T) {
	tests := []struct {
		inputStatus order.Status
		inputUserId uint64
		inputShopId uint64
	}{
		{
			inputStatus: order.StatusNotPay,
			inputUserId: 111,
			inputShopId: 111,
		},
	}
	for _, test := range tests {
		result := getCondition(test.inputStatus, test.inputUserId, test.inputShopId)
		if !assert.Equal(t, test.inputUserId, result.UserId) {
			return
		}
	}
}

func TestShopOrders(t *testing.T) {
	tests := []struct {
		inputQuery *QueryOrders
		wantCount  uint64
		err        error
	}{
		{
			inputQuery: &QueryOrders{
				OderId:   555555,
				ShopId:   555555,
				UserIds:  []uint64{555555},
				PageSize: 20,
				Page:     1,
				CreateTime: &TimeRange{
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
				Status: 11,
			},
			wantCount: 0,
			err:       nil,
		},
	}
	for _, test := range tests {
		_, count, err := ShopOrders(test.inputQuery)
		if !assert.Equal(t, test.err, err) {
			return
		}
		if !assert.Equal(t, count, test.wantCount) {
			return
		}
	}
}

func TestOrderStatusNums(t *testing.T) {
	tests := []struct {
		inputUserId uint64
		inputShopId uint64
		wantNum     int
	}{
		{
			inputUserId: 111,
			inputShopId: 111,
			wantNum:     0,
		},
	}
	for _, test := range tests {
		results, _ := OrderStatusNums(test.inputUserId, test.inputShopId)
		if !assert.Equal(t, len(results), test.wantNum) {
			return
		}
	}
}

func TestGetOrder(t *testing.T) {
	tests := []struct {
		inputOrderId     uint64
		inputInformation *order.Information
	}{
		{
			inputOrderId: 66666666,
			inputInformation: &order.Information{
				OrderId: 66666666,
				ShopId:  66666666,
				UserId:  66666666,
				Status:  order.StatusNotPay,
			},
		},
	}
	for _, test := range tests {
		database.GetDB(constants.DatabaseConfigKey).
			Model(&order.Information{}).Create(test.inputInformation)
		result, _ := GetOrder(test.inputOrderId)
		if !assert.Equal(t, test.inputOrderId, result.OrderId) {
			return
		}
	}
}

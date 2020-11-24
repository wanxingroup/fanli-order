package cancel

import (
	"os"
	"testing"

	"dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/utils/test"
)

func TestMain(m *testing.M) {

	test.Init()

	code := m.Run()

	test.Release()

	os.Exit(code)
}

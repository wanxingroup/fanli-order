package operation

import (
	"github.com/pkg/errors"
)

var ErrorOrderGoodsListEmpty = ErrorAction{
	Error:  errors.New("goods list empty"),
	Action: "check list",
}

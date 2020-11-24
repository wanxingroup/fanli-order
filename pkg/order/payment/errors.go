package payment

import (
	"errors"
)

var errorDuplicatedTransactionId = errors.New("transaction id was exists")

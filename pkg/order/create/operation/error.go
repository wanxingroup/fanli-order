package operation

import (
	"github.com/pkg/errors"
)

type ErrorAction struct {
	Action string
	Error  error
}

func (errorAction ErrorAction) WithError(err error) ErrorAction {

	errorAction.Error = err
	return errorAction
}

func (errorAction ErrorAction) WithAction(action string) ErrorAction {

	errorAction.Action = action
	return errorAction
}

func (errorAction ErrorAction) WrapError(err error) ErrorAction {

	errorAction.Error = errors.Wrap(errorAction.Error, err.Error())
	return errorAction
}

func (errorAction ErrorAction) Clone() *ErrorAction {

	return &errorAction
}

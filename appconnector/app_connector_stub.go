package appconnector

import (
	"errors"
)

var ErrorStub = errors.New("Stub")

type appConnectorStub struct {
}

func (s *appConnectorStub) ReturnFalse(connectInfo Connect) error {
	return ErrorStub
}

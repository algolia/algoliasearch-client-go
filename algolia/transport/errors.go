package transport

import (
	"errors"
	"net"
)

var (
	NoMoreHostToTryErr = errors.New("All hosts have been contacted unsuccessfully")
)

// NetError is used internally to differentiate regular error from errors
// following the net.Error interface in order to propagate them with a custom
// message.
type NetError struct {
	msg         string
	isTimeout   bool
	isTemporary bool
}

func NewNetError(err net.Error, msg string) *NetError {
	return &NetError{
		msg:         msg,
		isTimeout:   err.Timeout(),
		isTemporary: err.Temporary(),
	}
}

func (e *NetError) Error() string   { return e.msg }
func (e *NetError) Timeout() bool   { return e.isTimeout }
func (e *NetError) Temporary() bool { return e.isTemporary }

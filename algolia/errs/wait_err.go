package errs

type WaitError struct {
	msg string
}

func NewWaitError(msg string) *WaitError {
	return &WaitError{
		msg: msg,
	}
}

func (e *WaitError) Error() string {
	return e.msg
}

type WaitKeyUpdateError struct{}

func (e *WaitKeyUpdateError) Error() string {
	return "`apiKey` is required when waiting for an `update` operation."
}

type WaitKeyOperationError struct{}

func (e *WaitKeyOperationError) Error() string {
	return "`operation` must be one of `add`, `update` or `delete`."
}

package errs

type WaitError struct {
	msg string
}

func NewWaitError(msg string) *WaitError {
	return &WaitError{
		msg: msg,
	}
}

func (e WaitError) Error() string {
	return e.msg
}

func (e WaitError) Is(target error) bool {
	_, ok := target.(*WaitError)

	return ok
}

type WaitKeyUpdateError struct{}

func (e WaitKeyUpdateError) Error() string {
	return "`apiKey` is required when waiting for an `update` operation."
}

func (e WaitKeyUpdateError) Is(target error) bool {
	_, ok := target.(*WaitKeyUpdateError)

	return ok
}

type WaitKeyOperationError struct{}

func (e WaitKeyOperationError) Error() string {
	return "`operation` must be one of `add`, `update` or `delete`."
}

func (e WaitKeyOperationError) Is(target error) bool {
	_, ok := target.(*WaitKeyOperationError)

	return ok
}

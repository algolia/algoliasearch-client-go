package errs

var ErrNoMoreHostToTry = NewNoMoreHostToTryError()

type NoMoreHostToTryError struct {
	intermediateNetworkErrors []error
}

func NewNoMoreHostToTryError(errs ...error) *NoMoreHostToTryError {
	return &NoMoreHostToTryError{
		intermediateNetworkErrors: errs,
	}
}

func (e *NoMoreHostToTryError) IntermediateNetworkErrors() []error {
	return e.intermediateNetworkErrors
}

func (e *NoMoreHostToTryError) Error() string {
	return "all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. You can use opt.ExposeIntermediateNetworkErrors(true) to investigate."
}

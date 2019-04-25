package insights

type ClientInterface interface {
	User(userToken string) *UserClient
	SendEvent(event Event, opts ...interface{}) (res StatusMessageRes, err error)
	SendEvents(events []Event, opts ...interface{}) (res StatusMessageRes, err error)
}

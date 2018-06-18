package algoliasearch

import "fmt"

type analytics struct {
	client         *client
	abTestingRoute string
}

func NewAnalytics(client *client) *analytics {
	return &analytics{
		client:         client,
		abTestingRoute: "/2/abtests",
	}
}

func (a *analytics) AddABTest(abTest ABTest) (res ABTestTaskRes, err error) {
	path := a.abTestingRoute
	err = a.client.request(&res, "POST", path, abTest, analyticsCall, nil)
	return
}

func (a *analytics) StopABTest(id int) (res ABTestTaskRes, err error) {
	path := fmt.Sprintf("%s/%d/stop", a.abTestingRoute, id)
	err = a.client.request(&res, "POST", path, nil, analyticsCall, nil)
	return
}

func (a *analytics) DeleteABTest(id int) (res ABTestTaskRes, err error) {
	path := fmt.Sprintf("%s/%d", a.abTestingRoute, id)
	err = a.client.request(&res, "DELETE", path, nil, analyticsCall, nil)
	return
}

func (a *analytics) GetABTest(id int) (res ABTestResponse, err error) {
	path := fmt.Sprintf("%s/%d", a.abTestingRoute, id)
	err = a.client.request(&res, "GET", path, nil, analyticsCall, nil)
	return
}

func (a *analytics) GetABTests(params Map) (res GetABTestsRes, err error) {
	path := a.abTestingRoute
	err = a.client.request(&res, "GET", path, params, analyticsCall, nil)
	return
}

func (a *analytics) WaitTask(task ABTestTaskRes) (err error) {
	return a.client.WaitTask(task.Index, task.TaskID)
}

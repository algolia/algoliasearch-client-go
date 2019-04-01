package algolia

import (
	"sync"
)

type Waitable interface {
	Wait() error
}

type await struct {
	sync.Mutex
	waitables []Waitable
}

func Await() *await {
	return new(await)
}

func Wait(waitables ...Waitable) error {
	return Await().Wait(waitables...)
}

func (a *await) Collect(waitables ...Waitable) {
	a.Lock()
	a.waitables = append(a.waitables, waitables...)
	a.Unlock()
}

func (a *await) Wait() error {
	a.Lock()
	defer a.Unlock()

	var wg sync.WaitGroup
	errs := make(chan error, len(a.waitables))

	for _, w := range a.waitables {
		wg.Add(1)
		go func(wg *sync.WaitGroup, w Waitable, errs chan<- error) {
			errs <- w.Wait()
			wg.Done()
		}(&wg, w, errs)
	}

	wg.Wait()
	close(errs)
	a.waitables = nil

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

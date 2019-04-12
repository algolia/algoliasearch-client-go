package algolia

import (
	"sync"
)

// Waitable represents any asynchronous and potentially long-blocking operation
// which may eventually fail. Most of Algolia response objects are implementing
// the `algolia.Waitable` interface in order to wait for their completion using
// `algolia.Await()` or `algolia.Wait()` provided helpers.
type Waitable interface {
	Wait() error
}

type await struct {
	sync.Mutex
	waitables []Waitable
}

// Await returns a newly instantiated `await` object. This instance can collect
// objects implementing the `algolia.Waitable` interface, including most of the
// Algolia response objects, and wait for their completion in a concurrent
// fashion.
func Await() *await {
	return new(await)
}

// Wait blocks until all the given `algolia.Waitable` objects have completed. If
// one of the objects returned an error upon `Wait` invocation, this error is
// returned. Otherwise, nil is returned.
func Wait(waitables ...Waitable) error {
	await := Await()
	await.Collect(waitables...)
	return await.Wait()
}

// Collect holds references to the given `algolia.Waitable` objects in order to
// wait for their completion once the `Wait` method will be invoked. Calling
// `Collect` from multiple goroutines is safe.
func (a *await) Collect(waitables ...Waitable) {
	a.Lock()
	a.waitables = append(a.waitables, waitables...)
	a.Unlock()
}

// Wait blocks until all the collected `algolia.Waitable` objects have
// completed. If one of the objects returned an error upon `Wait` invocation,
// this error is returned. Otherwise, nil is returned. Calling `Wait` from
// multiple goroutines is safe.
//
// Upon successful completion, the `await` object can be reused directly to
// collect other `algolia.Waitable` objects.
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

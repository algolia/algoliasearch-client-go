package wait

// Waitable represents any asynchronous and potentially long-blocking operation
// which may eventually fail. Most of Algolia response objects are implementing
// the `algolia.Waitable` interface in order to wait for their completion using
// `algolia.NewGroup()` or `algolia.Wait()` provided helpers.
type Waitable interface {
	Wait(opts ...interface{}) error
}

package opt

import (
	"time"
)

const (
	maxDurationBetweenWaits = 10 * time.Second
)

// WaitConfigurationOption stores the logic to change the delay between waits
type WaitConfigurationOption struct {
	DelayGrowth func(*time.Duration) time.Duration
}

// DefaultWaitConfiguration returns the new delay to apply between waits
func DefaultWaitConfiguration() *WaitConfigurationOption {
	return &WaitConfigurationOption{DelayGrowth: func(d *time.Duration) time.Duration {
		if d == nil {
			return time.Second
		}
		res := *d * 2
		if res > maxDurationBetweenWaits {
			return maxDurationBetweenWaits
		}
		return res
	}}
}

// Get retrieves the actual value of the option parameter.
func (w *WaitConfigurationOption) Get() func(*time.Duration) time.Duration {
	if w == nil {
		return DefaultWaitConfiguration().DelayGrowth
	}
	return w.DelayGrowth
}

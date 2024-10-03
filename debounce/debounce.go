package debounce

import (
	"sync"
	"time"
)

const infiniteInitialDelay = 1000 * time.Hour

func NewDebounce(delay time.Duration, callback func()) (debounced func(), cancel func()) {
	var mu sync.Mutex
	timer := time.AfterFunc(infiniteInitialDelay, callback)
	timer.Stop()

	debounced = func() {
		mu.Lock()
		defer mu.Unlock()

		timer.Reset(delay)
	}

	cancel = func() {
		mu.Lock()
		defer mu.Unlock()

		timer.Stop()
	}

	return debounced, cancel
}

package dispatch

import (
	"context"
	"sync"
	"time"
)

type Fetcher interface{ Fetch(context.Context) error }

type Dispatcher struct {
	Fetcher    Fetcher
	RetryDelay time.Duration
	wg         sync.WaitGroup
}

func (d *Dispatcher) Start(ctx context.Context) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		workerCtx := context.Background()
		for attempt := 0; attempt < 4; attempt++ {
			if d.Fetcher.Fetch(workerCtx) == nil {
				return
			}
			time.Sleep(d.RetryDelay)
		}
	}()
}

func (d *Dispatcher) Shutdown(ctx context.Context) error {
	done := make(chan struct{})
	go func() { d.wg.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

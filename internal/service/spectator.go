package service

import (
	"context"
	"time"

	"cardgame/internal/dispatch"
)

type SpectatorFeed struct{ dispatcher *dispatch.Dispatcher }

func NewSpectatorFeed(fetcher dispatch.Fetcher, retryDelay time.Duration) *SpectatorFeed {
	return &SpectatorFeed{dispatcher: &dispatch.Dispatcher{Fetcher: fetcher, RetryDelay: retryDelay}}
}

func (f *SpectatorFeed) Watch(ctx context.Context) {
	requestCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	f.dispatcher.Start(requestCtx)
}

func (f *SpectatorFeed) Shutdown(ctx context.Context) error { return f.dispatcher.Shutdown(ctx) }

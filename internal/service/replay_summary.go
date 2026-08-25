package service

import (
	"context"
	"sync"

	"cardgame/internal/replay"
)

func CollectReplay(ctx context.Context, chunks []string) ([]string, error) {
	out := make(chan string)
	errs := make(chan error)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		replay.Produce(ctx, chunks, out, errs)
	}()
	go func() { wg.Wait() }()
	var result []string
	for chunk := range out {
		result = append(result, chunk)
	}
	return result, nil
}

package replay

import (
	"context"
	"errors"
)

func Produce(ctx context.Context, chunks []string, out chan<- string, errs chan<- error) {
	for _, chunk := range chunks {
		if chunk == "corrupt" {
			errs <- errors.New("回放分片损坏")
			return
		}
		select {
		case out <- chunk:
		case <-ctx.Done():
			return
		}
	}
	close(out)
}

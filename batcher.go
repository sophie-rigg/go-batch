package batcher

import (
	"context"
)

type Client[T any] struct {
	batchSize int
}

func New[T any](batchSize int) *Client[T] {
	return &Client[T]{
		batchSize: batchSize,
	}
}

func (c *Client[T]) Run(ctx context.Context, do func(ctx context.Context, object []T) error) (chan T, chan error) {
	channel := make(chan T, c.batchSize)
	errChannel := make(chan error)
	go func() {
		var (
			batch []T
			err   error
		)
		for object := range channel {
			batch = append(batch, object)
			if len(batch) >= c.batchSize {
				err = do(ctx, batch)
				if err != nil {
					errChannel <- err
				}
				batch = batch[:0]
			}
		}
		if len(batch) > 0 {
			err = do(ctx, batch)
			if err != nil {
				errChannel <- err
			}
		}
	}()

	return channel, errChannel
}

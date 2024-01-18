package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/sophie-rigg/go-batch"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	channel, errChannel := batcher.New[object](10).Run(ctx, func(ctx context.Context, object []object) error {
		var itemsInBatch []string
		for _, v := range object {
			itemsInBatch = append(itemsInBatch, v.val)
		}
		fmt.Println(itemsInBatch)
		return errors.New("error")
	})
	go func() {
		select {
		case err := <-errChannel:
			fmt.Println(err)
			cancelCtx()
		}
	}()
	for i := 0; i < 100; i++ {
		channel <- object{
			id:  i,
			val: fmt.Sprintf("val %d", i),
		}
	}
	close(channel)
	close(errChannel)
}

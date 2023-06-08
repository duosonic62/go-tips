package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// cancel
	ctx, cancel := context.WithCancel(context.Background())
	// timeout
	// ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	go f(ctx, &wg)

	time.Sleep(10 * time.Second)
	cancel() // キャンセル処理
	wg.Wait()
}

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("goroutine: doing")
		time.Sleep(1 * time.Second)
	}
}

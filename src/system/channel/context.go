package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel() // sub()が終了したことを通知
	}()
	<-ctx.Done() // sub()が終了するのを待つ
	fmt.Println("all tasks are finished")
}

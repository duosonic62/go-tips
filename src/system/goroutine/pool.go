package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manually added: 1")
	pool.Put("manually added: 2")

	// GCを実行すると、sync.Poolに入っている値が消える
	//runtime.GC()

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

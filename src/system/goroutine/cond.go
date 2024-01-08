package main

import (
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			println(name)
		}(name)
	}

	println("Ready?")
	time.Sleep(time.Second)
	println("Go!")
	cond.Broadcast()
	time.Sleep(3 * time.Second)
}

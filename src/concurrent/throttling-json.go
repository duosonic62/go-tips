package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	// download
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	// 並行数が20までに制限する
	limit := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			// limitに追加する
			limit <- struct{}{}
			defer wg.Done()

			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)

			// limitから取り出す
			<-limit
		}()
	}
	wg.Wait()

	fmt.Printf("%v\n", time.Since(before))
}

package main

import (
	"log"
	"net/http"
	// for get profiles
	_ "net/http/pprof"
	"sync"
)

func heavyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 100000; i++ {
		s = append(s, "magical pandas")
	}
}

// you can get profile http://localhost:8080/debug/pprof/profile
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go heavyFunc(&wg)
		wg.Wait()
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func stoppableGenerator(msg string, quit chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()

	return ch
}

func main() {
	quit := make(chan bool)
	ch := stoppableGenerator("Hi!", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}

	quit <- true
}

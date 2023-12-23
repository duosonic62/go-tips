package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := time.After(10 * time.Second)
	<-timerChan
	fmt.Println("time is over!")
}

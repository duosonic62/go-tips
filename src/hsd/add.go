package main

import "time"

func Add(a, b int) int {
	result := a + b

	// 重たい処理
	time.Sleep(3 * time.Second)
	return result
}

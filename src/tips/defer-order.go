package main

func main() {
	normalDefer()
	funcDefer()
}

func normalDefer() {
	n := 1
	// nはキャプチャされている
	defer println(n)

	n = 2
}

func funcDefer() {
	n := 1
	// nはキャプチャされない
	defer func() {
		println(n)
	}()

	n = 2
}

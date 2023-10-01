package main

func DoSomething(s string) {
	if len(s) > 100 {
		panic("something wrong")
	}
}

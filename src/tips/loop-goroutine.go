package main

import (
	"fmt"
	"sync"
)

type BigStruct struct {
	a int
	b string
}

func main() {
	var wg sync.WaitGroup
	bigstructs := []BigStruct{
		{
			a: 1,
			b: "johnny",
		},
		{
			a: 1,
			b: "smith",
		},
	}

	for _, bigstruct := range bigstructs {
		wg.Add(1)

		// 引数で渡す
		go func(bs *BigStruct) {
			defer wg.Done()
			fmt.Printf("%d %s \n", bs.a, bs.b)
		}(&bigstruct)
	}

	// インデックスでアクセス
	for i, _ := range bigstructs {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			fmt.Printf("%d %s \n", bigstructs[i].a, bigstructs[i].b)
		}()
	}
}

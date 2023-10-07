package benchmark

import "fmt"

func DoSomething() {

}

// メモリアロケーションの悪い例
func MakeSomethingBad(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%05d.", i)) // 毎回appendする
	}
	return r
}

// メモリアローけションの良い例
func MakeSomethingGood(n int) []string {
	r := make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%05d.", i)
	}
	return r
}

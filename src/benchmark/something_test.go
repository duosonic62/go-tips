package benchmark

import "testing"

// exec
// go test -bench . (or -bench DoSomething)
// calc memories
// go test -bench -benchmem .
func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}

func BenchmarkMakeSomethingBad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeSomethingBad(1000)
	}
}
func BenchmarkMakeSomethingGood(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeSomethingGood(1000)
	}
}

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

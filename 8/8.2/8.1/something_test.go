package something

import "testing"

func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}

// $ go test -benchmem -bench DoSomething
// goos: darwin
// goarch: amd64
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkDoSomething-12         1000000000               0.2711 ns/op          0 B/op          0 allocs/op
// PASS
// ok      /go-programming-essence/8/8.2/8.1      0.622s

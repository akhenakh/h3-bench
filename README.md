Benchmark to compare official [h3 Go](https://github.com/uber/h3-go) (using cgo), and a [transpiled C to Go](https://github.com/akhenakh/goh3) using [ccgo](https://pkg.go.dev/modernc.org/ccgo/v3#section-readme).

```
 go test -bench=. -benchtime=5s -benchmem
goos: linux
goarch: amd64
pkg: github.com/akhenakh/h3-bench
cpu: Intel(R) Core(TM) i5-4288U CPU @ 2.60GHz
BenchmarkToGeoCCGO-4            10657645               555.5 ns/op             0 B/op          0 allocs/op
BenchmarkToGeoCGO-4             10092243               561.7 ns/op            16 B/op          1 allocs/op
BenchmarkFromToCCGO-4            2142484              2872 ns/op               0 B/op          0 allocs/op
BenchmarkFromToCGO-4             1870581              3172 ns/op              32 B/op          2 allocs/op
PASS
ok      github.com/akhenakh/h3-bench    30.984s
```
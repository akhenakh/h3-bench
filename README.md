Benchmark to compare official [h3 Go](https://github.com/uber/h3-go) (using cgo), and a [transpiled C to Go](https://github.com/akhenakh/goh3) using [ccgo](https://pkg.go.dev/modernc.org/ccgo/v3#section-readme).

```
goos: linux
goarch: amd64
pkg: github.com/akhenakh/h3-bench
cpu: Intel(R) Core(TM) i5-4288U CPU @ 2.60GHz
BenchmarkToGeoCCGO-4            10657645               555.5 ns/op             0 B/op          0 allocs/op
BenchmarkToGeoCGO-4             10092243               561.7 ns/op            16 B/op          1 allocs/op
BenchmarkFromToCCGO-4            2142484              2872 ns/op               0 B/op          0 allocs/op
BenchmarkFromToCGO-4             1870581              3172 ns/op              32 B/op          2 allocs/op
PASS

goos: linux
goarch: amd64
pkg: github.com/akhenakh/h3-bench
cpu: Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz
BenchmarkToGeoCCGO-8    	32722182	       367.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkToGeoCGO-8     	30941000	       389.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkFromToCCGO-8   	 6580898	      1831 ns/op	       0 B/op	       0 allocs/op
BenchmarkFromToCGO-8    	 5373619	      2230 ns/op	      32 B/op	       2 allocs/op
PASS

goos: darwin
goarch: arm64
pkg: github.com/akhenakh/h3-bench
BenchmarkToGeoCCGO-8    	39700294	       275.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkToGeoCGO-8     	46501443	       257.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkFromToCCGO-8   	 8444726	      1420 ns/op	       0 B/op	       0 allocs/op
BenchmarkFromToCGO-8    	11856222	      1012 ns/op	      32 B/op	       2 allocs/op
```

Comparing geoToH3
```
Linux Intel
cat testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv | time ./h3-3.7.2/build/bin/geoToH3 -r 9  > /dev/null
~/Downloads/h3-3.7.2/build/bin/geoToH3 -r 9 > /dev/null  0.63s user 0.00s system 99% cpu 0.633 total
cat testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv | time ./cmd/geoToH3/geoToH3 -resolution 9  > /dev/null
./cmd/geoToH3/geoToH3 -resolution 9 > /dev/null  0.53s user 0.09s system 102% cpu 0.607 total

Darwin M1
cat testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv | time ../goh3/h3-3.7.2/build/bin/geoToH3 -r 9 > /dev/null
../goh3/h3-3.7.2/build/bin/geoToH3 -r 9 > /dev/null  0.31s user 0.00s system 99% cpu 0.309 total
cat testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv  testdata/populated.csv testdata/populated.csv testdata/populated.csv testdata/populated.csv | time ./cmd/geoToH3/geoToH3 -resolution 9 > /dev/null
./cmd/geoToH3/geoToH3 -resolution 9 > /dev/null  0.28s user 0.04s system 102% cpu 0.310 total

```

# Perf of `buildPutCall()`

Follow perf depending on refactoring commits:

1. f8b5a7 ✅ Bench buildCallPut()
2. d140be ♻️ Remove global vars used as temporary buffers
3. a77786 ⚡️ Remove an unnecessary memory allocation
4. cb561e ⚡️ Reuse []byte buffer to reduce memory allocation

These commits aims to remove ome global variables while keeping the same performance.

## Benchmark

    go test -benchmem -bench . ./pkg/rainbow/api
    goos: linux
    goarch: amd64
    pkg: github.com/teal-finance/rainbow/pkg/rainbow/api
    cpu: Intel(R) Core(TM) i7-3610QM CPU @ 2.30GHz
    BenchmarkAlign_buildCallPut/Psy-8      241720     5329 ns/op    2304 B/op    12 allocs/op
    BenchmarkAlign_buildCallPut/Deribit-8    1190  1015805 ns/op  510228 B/op  3118 allocs/op
    PASS
    ok      github.com/teal-finance/rainbow/pkg/rainbow/api 2.668s

## Table of results

Metrics from BenchmarkAlign_buildCallPut/Deribit-8:

| #  |     ns/op |    B/op | allocs/op
|----|:---------:|:-------:|:---------:
| 1. | 1 083 000 | 510 000 | 3118
| 2. | 1 280 000 | 544 650 | 5230
| 3. | 1 300 000 | 537 347 | 4770
| 4. | 1 170 000 | 510 200 | 3118

## Conclusion

The refactoring does not increase the number of memory allocations (3118 allocs/op),
but the final performance is not as good as the orignal code
based on global variables (`[]byte` buffers).

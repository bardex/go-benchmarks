# Golang Benchmarks

```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: benchmarks
cpu: Intel(R) Core(TM) i7-6700T CPU @ 2.80GHz

BenchmarkIndexedArrays/indexed_array-8              3258            360860 ns/op               0 B/op          0 allocs/op
BenchmarkIndexedArrays/indexed_slice-8              2397            487563 ns/op         1007620 B/op          1 allocs/op
BenchmarkIndexedArrays/indexed_map-8                  16          69703087 ns/op        24667450 B/op         20 allocs/op
BenchmarkIndexedArrays/indexed_map_without_capacity-8                  8         134207258 ns/op        54899052 B/op      38455 allocs/op

BenchmarkConcatOneLongString/concat-8                            1251682               951.3 ns/op          3104 B/op          9 allocs/op
BenchmarkConcatOneLongString/builder-8                           2184620               547.3 ns/op          1984 B/op          5 allocs/op

BenchmarkConcatManyShortStrings/concat-8                        10898780               108.1 ns/op            64 B/op          1 allocs/op
BenchmarkConcatManyShortStrings/builder-8                        7629721               204.7 ns/op           168 B/op          3 allocs/op
BenchmarkConcatManyShortStrings/join-8                          12462012                91.81 ns/op           64 B/op          1 allocs/op
BenchmarkConcatManyShortStrings/fmt-8                            2804323               434.6 ns/op           144 B/op          6 allocs/op

BenchmarkDefer/defer_loop-8                                      1000000              1028 ns/op             320 B/op         12 allocs/op
BenchmarkDefer/defer_loop_func-8                                 2631846               457.5 ns/op           160 B/op          2 allocs/op
BenchmarkDefer/no_defer-8                                        2882510               400.7 ns/op           160 B/op          2 allocs/op

BenchmarkEmptyInterface/empty_interface-8                        3878864               303.9 ns/op            72 B/op          6 allocs/op
BenchmarkEmptyInterface/concrete_type-8                          8076482               142.6 ns/op             8 B/op          2 allocs/op

BenchmarkInterfaces/interface-8                                      550           2151752 ns/op               0 B/op          0 allocs/op
BenchmarkInterfaces/concrete_type-8                                  631           1906573 ns/op               0 B/op          0 allocs/op

BenchmarkMaps50_50/mutex-8                                          2488            443200 ns/op           58485 B/op       1121 allocs/op
BenchmarkMaps50_50/rw_mutex-8                                       2506            430203 ns/op           58585 B/op       1122 allocs/op
BenchmarkMaps50_50/sync_map-8                                       2139            477192 ns/op           74022 B/op       1638 allocs/op

BenchmarkReflect/direct-8                                        3878752               304.2 ns/op            96 B/op          2 allocs/op
BenchmarkReflect/reflect-8                                        906993              1318 ns/op             352 B/op         13 allocs/op

BenchmarkReceiver/pointer_receiver-8                            1000000000               0.4803 ns/op          0 B/op          0 allocs/op
BenchmarkReceiver/regular_receiver-8                            1000000000               0.4101 ns/op          0 B/op          0 allocs/op

PASS
ok      benchmarks      32.611s
```
# go-benchmark

### String Concatenation
```bash
$ cd concatenation && go test -bench=. -benchmem .
goos: darwin
goarch: amd64
pkg: github.com/andrey1s/go-benchmark/concatenation
BenchmarkSfmt-4           	  100000	     12207 ns/op	    4705 B/op	     102 allocs/op
BenchmarkIfmt-4           	  100000	     12175 ns/op	    4289 B/op	     102 allocs/op
BenchmarkFfmt-4           	  100000	     23439 ns/op	    4289 B/op	     102 allocs/op
BenchmarkBuff-4           	  500000	      2516 ns/op	    6928 B/op	       7 allocs/op
BenchmarkJoin-4           	 1000000	      1041 ns/op	    1792 B/op	       1 allocs/op
BenchmarkJoinInt-4        	  300000	      5914 ns/op	    5248 B/op	     101 allocs/op
BenchmarkConc100-4        	 1000000	      1367 ns/op	    1792 B/op	       1 allocs/op
BenchmarkConc100Int-4     	  200000	      6145 ns/op	    5248 B/op	     101 allocs/op
BenchmarkConc100Float-4   	  100000	     19105 ns/op	    8448 B/op	     201 allocs/op
PASS
ok  	github.com/andrey1s/go-benchmark/concatenation	14.232s
```

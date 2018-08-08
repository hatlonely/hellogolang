
bench:
	go test -benchmem -cpuprofile cpu.out -memprofile mem.out -run=^$ github.com/hatlonely/hellogolang/benchmark/bench -bench ^Benchmark$
	go tool pprof -svg bench.test cpu.out > cpu.svg
	go tool pprof -svg bench.test mem.out > mem.svg

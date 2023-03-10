# 性能测试

```shell
go test -bench=. *


go test -bench=BenchmarkMarshalStdJson -run=BenchmarkMarshalStdJson -benchmem -cpuprofile stdjson-marshal.cpu
go tool pprof -svg ./json.test stdjson-marshal.cpu > stdjson-marshal.svg


go test -bench=BenchmarkMarshalJsonIterator -run=BenchmarkMarshalJsonIterator -benchmem -cpuprofile jsoniter-marshal.cpu
go tool pprof -svg ./json.test jsoniter-marshal.cpu > jsoniter-marshal.svg


go test -bench=BenchmarkMarshalEasyjson -run=BenchmarkMarshalEasyjson -benchmem -cpuprofile easyjson-marshal.cpu
go tool pprof -svg ./json.test easyjson-marshal.cpu > easyjson-marshal.svg


go test -bench=BenchmarkUnMarshalStdJson -run=BenchmarkUnMarshalStdJson -benchmem -cpuprofile stdjson-unmarshal.cpu
go tool pprof -svg ./json.test stdjson-unmarshal.cpu > stdjson-unmarshal.svg


go test -bench=BenchmarkUnMarshalJsonIterator -run=BenchmarkUnMarshalJsonIterator -benchmem -cpuprofile jsoniter-unmarshal.cpu
go tool pprof -svg ./json.test jsoniter-unmarshal.cpu > jsoniter-unmarshal.svg


go test -bench=BenchmarkUnMarshalEasyjson -run=BenchmarkUnMarshalEasyjson -benchmem -cpuprofile easyjson-unmarshal.cpu
go tool pprof -svg ./json.test easyjson-unmarshal.cpu > easyjson-unmarshal.svg
```
codegen: counter.pb.go

counter.pb.go: counter.proto
	protoc --go_out=plugins=grpc:. counter.proto

clean:
	-rm -rf counter.pb.go


proto: gogoprotobuf
	protoc --gofast_out=plugins=grpc:. creative.proto && mkdir gogoproto && mv creative.pb.go gogoproto
	protoc --go_out=plugins=grpc:. creative.proto && mkdir protobuf && mv creative.pb.go protobuf

gogoprotobuf:
	go get github.com/gogo/protobuf/protoc-gen-gofast
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/jsonpb
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto
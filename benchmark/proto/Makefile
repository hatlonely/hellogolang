
proto: gogoprotobuf
	protoc --gofast_out=plugins=grpc:. creative.proto && mkdir -p gogoproto && mv module/creative.pb.go gogoproto
	protoc --go_out=. --go_opt=paths=source_relative creative.proto && mkdir -p protobuf && mv creative.pb.go protobuf

gogoprotobuf:
	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/jsonpb
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto

protoc:
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

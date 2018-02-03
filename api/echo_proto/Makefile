codegen: echo.pb.go

echo.pb.go: echo.proto
	protoc --go_out=plugins=grpc:. echo.proto

clean:
	-rm -rf echo.pb.go

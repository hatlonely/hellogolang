codegen: gen-go

gen-go: echo.thrift
	thrift -r --gen go echo.thrift

clean:
	-rm -rf gen-go

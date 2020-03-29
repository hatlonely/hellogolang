export PATH:=${PATH}:${GOPATH}/bin
export PATH:=$(shell pwd)/third/go/bin:${PATH}
export PATH:=$(shell pwd)/third/protobuf/bin:${PATH}
export PATH:=$(shell pwd)/third/thrift/bin:${PATH}
export PATH:=$(shell pwd)/third/bison/bin:${PATH}

output: cmd/*/*.go Makefile api/counter_proto/counter.pb.go api/echo_proto/echo.pb.go api/echo_thrift/gen-go vendor
	@echo "compile"
	rm -rf output
	mkdir -p output/grpc/counter/bin 
	go build cmd/grpc/counter_client.go && mv main output/grpc/counter/bin/client
	go build cmd/grpc/counter_server.go && mv main output/grpc/counter/bin/server
	mkdir -p output/grpc/echo/bin
	go build cmd/grpc/echo_client.go && mv main output/grpc/echo/bin/client
	go build cmd/grpc/echo_server.go && mv main output/grpc/echo/bin/server
	mkdir -p output/thrift/echo/bin
	go build cmd/thrift/echo_client.go && mv main output/thrift/echo/bin/client
	go build cmd/thrift/echo_server.go && mv main output/thrift/echo/bin/server

vendor: go.mod
	@echo "install golang dependency"
	go mod vendor

%.pb.go: %.proto
	$(shell pwd)/third/protobuf/bin/protoc --go_out=plugins=grpc:. $<

.PHONY: protoc
protoc:
	@hash protoc 2>/dev/null || { \
		echo "install protobuf codegen tool protoc" && \
		mkdir -p third && cd third && \
		wget https://github.com/google/protobuf/releases/download/v3.2.0/protobuf-cpp-3.2.0.tar.gz && \
		tar -xzvf protobuf-cpp-3.2.0.tar.gz && \
		cd protobuf-3.2.0 && \
		./configure --prefix=`pwd`/../protobuf && \
		make -j8 && \
		make install && \
		cd ../.. && \
		protoc --version; \
	}
	@hash protoc-gen-go 2>/dev/null || { \
		echo "install protobuf golang plugin protoc-gen-go" && \
		go get -u github.com/golang/protobuf/{proto,protoc-gen-go}; \
	}

.PHONY: thrift
thrift: bison
	hash thrift 2>/dev/null || { \
		echo "install thrift" \
		mkdir -p third && cd third && \
		wget http://www-us.apache.org/dist/thrift/0.11.0/thrift-0.11.0.tar.gz && \
		tar -xzvf thrift-0.11.0.tar.gz && \
		cd thrift-0.11.0 && \
		./configure --prefix=`pwd`/../thrift \
			--with-qt4=no \
			--with-qt5=no \
			--with-c_glib=no \
			--with-csharp=no \
			--with-java=no \
			--with-erlang=no \
			--with-nodejs=no \
			--with-lua=no \
			--with-python=no \
			--with-perl=no \
			--with-php=no \
			--with-php_extension=no \
			--with-dart=no \
			--with-ruby=no \
			--with-haskell=no \
			--with-rs=no \
			--with-haxe=no \
			--with-dotnetcore=no \
			--with-d=no \
			--disable-shared \
			--disable-tests \
			--disable-tutorial && \
		make -j8 && \
		make install && \
		cd ../.. && \
		thrift --version; \
	}

.PHONY: bison
bison:
	hash bison 2>/dev/null && [[ "`bison --version | head -1 | awk '{print $$NF}'`" > "2.5" ]] || { \
		echo "install bison" \
		mkdir -p third && cd third && \
		wget http://mirrors.nju.edu.cn/gnu/bison/bison-3.0.5.tar.gz && \
		tar -xzvf bison-3.0.5.tar.gz && \
		cd bison-3.0.5 && \
		./configure --prefix=`pwd`/../bison && \
		make -j8 && \
		make install && \
		cd ../.. && \
		bison --version; \
	}

export PATH:=${PATH}:${GOPATH}/bin:$(shell pwd)/third/go/bin:$(shell pwd)/third/protobuf/bin:$(shell pwd)/third/cloc-1.76:$(shell pwd)/third/redis-3.2.8/src

.PHONY: all
all: build third vendor test stat

build: cmd/*/*.go internal/*/*.go pkg/*/*.go scripts/version.sh Makefile vendor
	@echo "compile"
	@go build -ldflags "-X 'main.AppVersion=`sh scripts/version.sh`'" cmd/server/main.go && \
	mkdir -p build/server/bin && mv main build/server/bin/server && \
	mkdir -p build/server/configs && cp configs/server/* build/server/configs && \
	go build -ldflags "-X 'main.AppVersion=`sh scripts/version.sh`'" cmd/client/main.go && \
	mkdir -p build/client/bin && mv main build/client/bin/client

vendor: glide.lock glide.yaml
	@echo "install golang dependency"
	glide install

.PHONY: api
api:
	@echo "generate protocol code"
	@thrift -r --gen go -o api/addservice api/addservice/add.thrift

.PHONY: test
test: vendor
	@echo "Run unit tests"
	- cd internal && go test -cover ./...
	- cd pkg && go test -cover ./...

.PHONY: stat
stat: cloc gocyclo
	@echo "code statistics"
	@cloc pkg Makefile --by-file
	@echo "circle complexity statistics"
	@gocyclo pkg
	@gocyclo pkg | awk '{sum+=$$1}END{printf("complexity: %s", sum)}'

.PHONY: clean
clean:
	rm -rf build

.PHONY: deep_clean
deep_clean:
	rm -rf build vendor third

third: protoc glide golang cloc gocyclo easyjson

.PHONY: protoc
protoc: golang
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

.PHONY: glide
glide: golang
	@mkdir -p $$GOPATH/bin
	@hash glide 2>/dev/null || { \
		echo "install glide" && \
		curl https://glide.sh/get | sh; \
	}

.PHONY: golang
golang:
	@hash go 2>/dev/null || { \
		echo "install go1.9" && \
		mkdir -p third && cd third && \
		wget https://dl.google.com/go/go1.9.linux-amd64.tar.gz && \
    	tar -xzvf go1.9.linux-amd64.tar.gz && \
		cd .. && \
		go version; \
	}

.PHONY: cloc
cloc:
	@hash cloc 2>/dev/null || { \
		echo "install cloc" && \
		mkdir -p third && cd third && \
		wget https://github.com/AlDanial/cloc/archive/v1.76.zip && \
		unzip v1.76.zip; \
	}

.PHONY: gocyclo
gocyclo: golang
	@hash gocyclo 2>/dev/null || { \
		echo "install gocyclo" && \
		go get -u github.com/fzipp/gocyclo; \
	}

.PHONY: easyjson
easyjson: golang
	@hash easyjson 2>/dev/null || { \
		echo "install easyjson" && \
		go get -u github.com/mailru/easyjson/...; \
	}

.PHONY: redis
redis:
	@hash redis 2>/dev/null || { \
		echo "install redis" && \
		mkdir -p third && cd third && \
		wget http://download.redis.io/releases/redis-3.2.8.tar.gz && \
		tar -xzvf redis-3.2.8.tar.gz && \
		cd redis-3.2.8 && \
		make -j8; \
	}

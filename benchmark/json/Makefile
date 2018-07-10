export PATH:=${PATH}:${GOPATH}/bin

all: ffjson easyjson json

json: book_ffjson.go book_easyjson.go

%_ffjson.go: %.go
	ffjson $<

%_easyjson.go: %.go
	easyjson $<

.PHONY: ffjson
ffjson:
	@hash ffjson 2>/dev/null || { \
		echo "install ffjson" && \
		go get -u github.com/pquerna/ffjson; \
	}

.PHONY: easyjson
easyjson:
	@hash easyjson 2>/dev/null || { \
		echo "install easyjson" && \
		go get -u github.com/mailru/easyjson/...; \
	}
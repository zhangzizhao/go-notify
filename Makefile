all: install

GOPATH:=$(CURDIR)
export GOPATH

fmt:
	gofmt -l -w -s src/

dep:fmt
	go get github.com/gin-gonic/gin
	go get git.apache.org/thrift.git/lib/go/thrift

build:
	cd src/notify && go build

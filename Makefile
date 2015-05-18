all: install

GOPATH:=$(CURDIR)
export GOPATH

fmt:
	gofmt -l -w -s src/

dep:fmt
	go get github.com/gin-gonic/gin

build:
	cd src/notify && go build

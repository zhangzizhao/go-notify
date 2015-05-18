all: install

GOPATH:=$(CURDIR)
export GOPATH

fmt:
	gofmt -l -w -s src/

dep:fmt
	go get github.com/gin-gonic/gin

run:
	cd src/notify && $(CURDIR)/bin/bee run --gendoc=true --downdoc=true

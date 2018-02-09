GOPATH := $(shell pwd)
.PHONY: clean

all:
	@GOPATH=$(GOPATH) go install telegram_login

clean:
	@rm -fr bin pkg
 
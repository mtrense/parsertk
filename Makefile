
# Makefile for parsertk

VERSION ?= 0.0.0

test:

test-prerequisites: install-tools

install-tools:
	go get github.com/onsi/ginkgo/ginkgo
	
### TEST ####################################################################

test-parsertk:
	ginkgo -r
test-parsertk-watch:
	ginkgo watch
test: test-parsertk
.PHONY: test-parsertk
.PHONY: test

clean:
	rm -r bin/* dist/*

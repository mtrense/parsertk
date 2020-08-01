
# Makefile for parsertk

VERSION ?= 0.0.0

test:

test-prerequisites:

install-tools:

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

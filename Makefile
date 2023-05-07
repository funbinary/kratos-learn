
.PHONY: build
# build
build:
	mkdir -p bin/
	go build  -o ./bin/ config/file/main.go

.PHONY: all
# generate all
all:
	make build;


.DEFAULT_GOAL := all

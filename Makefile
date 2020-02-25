.PHONY: build watch

build:
	go build .

tidy: build
	go fmt ./...
	go mod tidy

watch:
	modd

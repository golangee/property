all:
	go build
	golangci-lint run --enable-all --exclude-use-default=false --disable=paralleltest
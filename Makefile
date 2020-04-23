all: build
.PHONY: clean

tidy:
	GOPROXY=https://goproxy.io GO111MODULE=on go mod tidy

build:
	GOPROXY=https://goproxy.io GO111MODULE=on go build -v

clean:
	rm -rf myproject
	go clean -i .
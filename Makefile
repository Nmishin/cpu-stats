.PHONY: build build-x64 build-arm64

BINARY_NAME := cpu-stats-exporter

build: build-x64 build-arm64

build-x64:
	@echo "Building for x64"
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-x64 main.go

build-arm64:
	@echo "Building for arm64"
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-arm64 main.go

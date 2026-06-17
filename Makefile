.PHONY: run build test clean

APP_NAME=myapp
BUILD_DIR=./build

run:
	go run ./cmd/app/main.go

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/app/main.go

test:
	go test ./... -v

clean:
	rm -rf $(BUILD_DIR)

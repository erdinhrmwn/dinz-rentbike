.PHONY: run build test clean fmt

APP_NAME=dinz-rentbike
BUILD_DIR=./bin

run:
	go run ./cmd/app/main.go

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/app/main.go

test:
	go test ./... -v

fmt:
	goimports -w .
	goimports-reviser -rm-unused -project-name $(APP_NAME) -separate-named -set-alias -format -recursive .
	go fmt ./...

clean:
	rm -rf $(BUILD_DIR)

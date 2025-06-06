APP_NAME = oatcoin
BUILD_DIR = build
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build run test clean fmt lint

all: build

build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd

run: build
	@echo "🚀 Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

fmt:
	@echo "🧼 Formatting code..."
	go fmt ./...

lint:
	@echo "🔍 Linting code..."
	golangci-lint run

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR)

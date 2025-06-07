APP_NAME = oatcoin
ROOT_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST)))/..)
BUILD_DIR := $(ROOT_DIR)/build
GOFILES := $(shell find $(ROOT_DIR) -type f -name '*.go' -not -path "$(ROOT_DIR)/vendor/*")

.PHONY: all build clean fmt

all: build

build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	cd $(ROOT_DIR) && go build -o build/$(APP_NAME) ./cmd

fmt:
	@echo "🧼 Formatting code..."
	cd $(ROOT_DIR) && go fmt ./...

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR)

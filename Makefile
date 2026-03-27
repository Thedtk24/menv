VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags "-X main.version=$(VERSION)"
APP_NAME := menv
BUILD_DIR := build
BIN_NAME := $(BUILD_DIR)/$(APP_NAME)

.PHONY: all build install test clean

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BIN_NAME) main.go

install: build
	@echo "Installing $(APP_NAME) to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	cp $(BIN_NAME) ~/.local/bin/$(APP_NAME)
	@echo "Installation successful."

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

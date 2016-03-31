UNAME := $(shell sh -c 'uname')
VERSION := $(shell sh -c 'git describe --always --tags')
PATH := $(subst :,/bin:,$(GOPATH))/bin:$(PATH)

NAME=mochi
MAIN=./cmd/mochi/main.go
OUT=./build/$(NAME)

BUILD_DIR="./build"
BUILD_NUMBER=$(shell date "+%Y%m%d-%H%M%S")

# We have build folder which is same name with target
# http://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date
.PHONY: build

default: build

prepare:
	go get ./...
	mkdir -p $(BUILD_DIR)/$(BUILD_NUMBER)

build: prepare
	go build -o $(BUILD_DIR)/$(BUILD_NUMBER)/$(NAME) \
			-ldflags \
				"-X main.VERSION=$(VERSION)" \
		$(MAIN)
	ln -sf ./$(BUILD_NUMBER)/$(NAME) $(OUT)

clean:
	rm -rf build/*

repl:
	$(OUT)

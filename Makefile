# export GOPROXY=http://goproxy.cn
GO := GO111MODULE=on go
CUR_PWD := $(shell pwd)
BINARY_NAME = gintest
PROJECT_URL = ""
CONFIG_NAME = $(CUR_PWD)/config.yaml
COMMIT_SHA1 := $(shell git rev-parse HEAD )
BUILD_TIME := $(shell date +%Y-%m-%d\ %H:%M:%S)

LD_FLAGS=" -X '$(BINARY_NAME)/config.PROJECT_NAME=$(BINARY_NAME)' \
	-X '$(BINARY_NAME)/config.PROJECT_URL=$(PROJECT_URL)' \
	-X '$(BINARY_NAME)/config.COMMIT_SHA1=$(COMMIT_SHA1)' \
	-X '$(BINARY_NAME)/config.BUILD_TIME=$(BUILD_TIME)'"

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default: dev

clean:
	@echo ">>> Start cleaning"
	@rm -rf ./$(BINARY_NAME)*
	@rm -rf ./nohup.out
	@echo "<<< Cleaning is complete"

race:
	@echo ">>> Start detecting race conditions"
	@$(GO) build -o $(BINARY_NAME) -ldflags $(LD_FLAGS) -v -race
	@echo "<<< Race conditions detecting is complete"

gofmt:
	@echo ">>> Start code formatting"
	gofmt -s -w ${GOFILES}
	@echo "<<< Code formatting is complete"

cover:
	@echo ">>> Start code coverage testing"
	@$(GO) test -coverpkg="./..." -ldflags  $(LD_FLAGS)  -c -cover -covermode=atomic -o $(BINARY_NAME)
	@echo "<<< Code coverage testing is complete"

dev: clean main.go go.sum go.mod
	@echo ">>> [dev] Start building"
	@cp $(CUR_PWD)/config/config.dev.yaml $(CONFIG_NAME)
	@$(GO) build -o $(BINARY_NAME) -ldflags $(LD_FLAGS)
	@echo "<<< [dev] $(BINARY_NAME) build success"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit_SHA1: $(COMMIT_SHA1)"

test: clean main.go go.sum go.mod
	@echo ">>> [test] Start building"
	@cp $(CUR_PWD)/config/config.test.yaml $(CONFIG_NAME)
	@$(GO) build -o $(BINARY_NAME) -ldflags $(LD_FLAGS)
	@echo "<<< [test] $(BINARY_NAME) build success"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit_SHA1: $(COMMIT_SHA1)"

release: clean main.go go.sum go.mod
	@echo ">>> [release] Start building"
	@cp $(CUR_PWD)/config/config.release.yaml $(CONFIG_NAME)
	@$(GO) build -o -w -s $(BINARY_NAME) -ldflags $(LD_FLAGS) -gcflags "-N" -o
	@echo "<<< [release] $(BINARY_NAME) build success"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit_SHA1: $(COMMIT_SHA1)"

run:
	./$(BINARY_NAME) -c $(CONFIG_NAME)

prod:
	@nohup ./$(BINARY_NAME) -c $(CONFIG_NAME) -m release

.PHONY: clean dev test release run prod race gofmt cover
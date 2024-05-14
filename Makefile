PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
VERSION ::= $(shell git describe --always --tags --dirty)
BUILD_TIME ::= $(shell date "+%Y-%m-%d_%H:%M:%S%:z")
COMMITISH ::= $(shell git describe --always --dirty)
ifeq (${GOPATH},)
	GOPATH := ${HOME}/go
endif
OS ::= $(shell uname -s)
SHELL = /bin/bash
GO_VERSION=1.22

define build_binary
	@echo "Building $(1)/$(2)..."
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build \
		-a \
		-ldflags "-X github.com/cbuschka/tfvm/internal/build.buildInfoVersion=${VERSION} \
			-X github.com/cbuschka/tfvm/internal/build.buildInfoBuildTime=${BUILD_TIME} \
			-X github.com/cbuschka/tfvm/internal/build.buildInfoCommitish=${COMMITISH} \
			-X github.com/cbuschka/tfvm/internal/build.buildInfoOs=$(1) \
			-X github.com/cbuschka/tfvm/internal/build.buildInfoArch=$(2) \
			-extldflags \"-static\"" \
			-o dist/tfvm-$(1)_$(2)$(3) \
			cmd/tfvm/tfvm.go
endef

all:	clean build_linux build_windows build_macosx integration_test

check_go_version:
	@if [[ ! "$$(go version)" =~ ^.*go${GO_VERSION}.*$$ ]]; then \
		echo "Wrong go version. Expected go ${GO_VERSION}."; \
		exit 1; \
	else \
		echo "Go version ok."; \
	fi

init:	check_go_version
	@if [ "${OS}" != "Linux" ] && [ "${OS}" != "Darwin" ]; then \
		echo "Sorry only Linux and macOS supported as build platform. (This is ${OS}.)"; \
		exit 1; \
	fi; \
	mkdir -p ${GOPATH}

lint:	init
	@echo "### Running linter..."
	go install golang.org/x/lint/golint@latest
	${GOPATH}/bin/golint ./internal/... ./cmd/...

build:	test lint
	go vet ./...
	mkdir -p dist/

build_all:	build build_linux build_windows build_macosx

build_linux:	build_linux_amd64 build_linux_386 build_linux_arm64
	@echo "### Building Linux variants..."
	$(call build_binary,linux,amd64,)

build_linux_arm64:	build
	$(call build_binary,linux,arm64,)

build_linux_amd64:	build
	$(call build_binary,linux,amd64,)

build_linux_386:	build
	$(call build_binary,linux,386,)

.PHONY: clean
clean:	init
	@echo "### Cleaning up..."
	rm -rf ${PROJECT_DIR}/dist/ ${PROJECT_DIR}/.cache/

format:
	gofmt -l -w -s ${PROJECT_DIR}

build_windows:	build
	@echo "### Building Windows variants..."
	$(call build_binary,windows,amd64,.exe)
	$(call build_binary,windows,386,.exe)

build_macosx:	build
	@echo "### Building macOS variants..."
	$(call build_binary,darwin,amd64,)
	$(call build_binary,darwin,arm64,)

test:	init
	@echo "### Running unit tests..."
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./internal/... ./cmd/...

integration_test:
	@echo "### Running integration tests..."
	${PROJECT_DIR}/scripts/run-integration-tests.sh

dump_tf_releases:	init
	(echo -e "package state\n\nconst defaultStateJSON = \`" && \
	go run ./cmd/dump_tf_releases/ ./internal/... && \
	echo '`' ) > ${PROJECT_DIR}/internal/inventory/default_inventory_state.go.new && \
	mv ${PROJECT_DIR}/internal/inventory/default_inventory_state.go.new ${PROJECT_DIR}/internal/inventory/state/default_inventory_state.go

build_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:${GO_VERSION}-buster \
		make build

test_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:${GO_VERSION} \
		make test

update_dependencies:	init
	go get -t -u ./...

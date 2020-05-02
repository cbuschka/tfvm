PROJECT_DIR ::= ${PWD}
VERSION ::= $(shell git describe --always --tags --dirty)
BUILD_TIME ::= $(shell date "+%Y-%m-%d_%H:%M:%S%:z")
COMMITISH ::= $(shell git describe --always --dirty)

define build_binary
	echo "Building $(1)/$(2)..."
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -X github.com/cbuschka/tfvm.buildInfoOs=$(1) -X github.com/cbuschka/tfvm.buildInfoArch=$(2) -extldflags \"-static\"" -o dist/tfvm-$(1)_$(2)$(3) cmd/main.go
endef

all:	clean build_linux build_windows build_macosx

lint:
	go get -u golang.org/x/lint/golint
	golint ./... 

build:	test lint
	go vet ./...
	mkdir -p dist/

build_linux:	build
	$(call build_binary,linux,amd64,)
	$(call build_binary,linux,386,)

clean:
	rm -rf ${PROJECT_DIR}/dist/ ${PROJECT_DIR}/.cache/

format:
	go fmt ./...

build_windows:	build
	$(call build_binary,windows,amd64,.exe)
	$(call build_binary,windows,386,.exe)

build_macosx:	build
	$(call build_binary,darwin,amd64,)

test:
	go test ./...

build_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:1.13-stretch \
		make build

test_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:1.13-stretch \
		make test

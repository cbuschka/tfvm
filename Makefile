PROJECT_DIR ::= ${PWD}
VERSION ::= $(shell git describe --always --tags --dirty)
BUILD_TIME ::= $(shell date "+%Y-%m-%d_%H:%M:%S%:z")
COMMITISH ::= $(shell git describe --always --dirty)

all:	clean test lint build build_windows build_macosx

lint:
	go get -u golang.org/x/lint/golint
	golint ./... 

build:
	go vet ./...
	mkdir -p dist/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -X github.com/cbuschka/tfvm.buildInfoOs=linux -X github.com/cbuschka/tfvm.buildInfoArch=amd64 -extldflags \"-static\"" -o dist/tfvm-linux_amd64 cmd/main.go

clean:
	rm -rf ${PROJECT_DIR}/dist/ ${PROJECT_DIR}/.cache/

format:
	go fmt ./...

build_windows:
	mkdir -p dist/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -X github.com/cbuschka/tfvm.buildInfoOs=windows -X github.com/cbuschka/tfvm.buildInfoArch=amd64 -extldflags \"-static\"" -o dist/tfvm-windows_amd64.exe cmd/main.go

build_macosx:
	mkdir -p dist/
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -X github.com/cbuschka/tfvm.buildInfoOs=darwin -X github.com/cbuschka/tfvm.buildInfoArch=amd64 -extldflags \"-static\"" -o dist/tfvm-darwin_amd64 cmd/main.go

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

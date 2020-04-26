PROJECT_DIR ::= ${PWD}
VERSION ::= $(shell git describe --always --tags --dirty)
BUILD_TIME ::= $(shell date "+%Y-%m-%d_%H:%M:%S%:z")
COMMITISH ::= $(shell git describe --always --dirty)

all:	test build lint build_windows_and_macosx

lint:
	go get -u golang.org/x/lint/golint
	golint ./... 

build:
	go vet ./...
	mkdir -p dist/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -extldflags \"-static\"" -o dist/tfvm-linux_amd64 cmd/main.go

format:
	go fmt ./...

build_windows_and_macosx:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -extldflags \"-static\"" -o dist/tfvm-windows_amd64 cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags "-X github.com/cbuschka/tfvm.buildInfoVersion=${VERSION} -X github.com/cbuschka/tfvm.buildInfoBuildTime=${BUILD_TIME} -X github.com/cbuschka/tfvm.buildInfoCommitish=${COMMITISH} -extldflags \"-static\"" -o dist/tfvm-darwin_amd64 cmd/main.go

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

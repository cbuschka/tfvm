PROJECT_DIR ::= ${PWD}

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags \"-static\"' -o tfvm cmd/main.go

test:
	go test ./...

all:	test build

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

PROJECT_DIR ::= ${PWD}

build:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		golang:1.13-stretch \
		/bin/bash -c "HOME=/build GO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags \"-static\"' -o tfvm cmd/main.go"

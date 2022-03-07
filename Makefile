SOURCE = ./...
OS = linux
ARCH = amd64

vet:
	go vet $(SOURCE)

test-fmt:
	test -z $(shell go fmt $(SOURCE))

test: vet test-fmt
	go test -cover ./... -count 1

build:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) \
    go build -a -ldflags '-extldflags "-static"' \
    -o ensure-tfenv-versions .
SOURCE = ./...

tidy:
	go mod tidy

generate:
	go generate

vet:
	go vet $(SOURCE)

test-fmt:
	test -z $(shell go fmt $(SOURCE))

test: generate vet test-fmt
	go test -cover $(SOURCE) -count 1

build:
	go run github.com/goreleaser/goreleaser build --snapshot --rm-dist
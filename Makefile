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

integration-test: tidy generate
	go run main.go -v

build:
	go run github.com/goreleaser/goreleaser build --snapshot --rm-dist

release:
	go run github.com/goreleaser/goreleaser release --rm-dist
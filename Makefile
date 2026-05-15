APP     := pulumiGo
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -X main.Version=$(VERSION)

.PHONY: build run clean test vet fmt tidy

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP) .

run:
	go run -ldflags "$(LDFLAGS)" . $(ARGS)

test:
	go test -race -count=1 ./...

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out | tail -1

vet:
	go vet ./...

fmt:
	gofmt -s -w .
	goimports -w . 2>/dev/null || true

tidy:
	go mod tidy

clean:
	rm -rf bin/

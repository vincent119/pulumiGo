APP     := pulumiGo
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -X main.Version=$(VERSION)

.PHONY: build run clean test vet fmt tidy

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP) .

run:
	go run -ldflags "$(LDFLAGS)" . $(ARGS)

test:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .
	goimports -w . 2>/dev/null || true

tidy:
	go mod tidy

clean:
	rm -rf bin/

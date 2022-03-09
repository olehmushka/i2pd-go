default: run

install:
	go mod download

test:
	go test ./...

test-force:
	go clean -testcache && go test ./...

coverage:
	go test -cover ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

run:
	go run main.go

lint:
	golangci-lint run --build-tags test ./...

build:
	@go build -o bin/social cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v ./...

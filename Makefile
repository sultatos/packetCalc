# Simple Makefile for a Go project

# Build the application
all: build test

templ: templ-install
	@templ generate

build: templ
	@go build -o ./bin/api cmd/api/main.go
# Run the application
run: templ
	@go run cmd/api/main.go
# Test the application
test:
	@go test ./... -v
# Clean the binary
clean:
	@rm -f ./bin/api
templ-install:
	which templ || go install github.com/a-h/templ/cmd/templ@latest

image: build test
	@docker build -t go-api .

docker-run:
	@docker run -p 8080:8080  go-api

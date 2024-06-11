build:
	@go build -o bin/ecom-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-api
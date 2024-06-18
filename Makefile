build:
	@go build -o bin/ecom-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-api


migration:
	@migrate create -ext sql -dir cmd/migrate/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
build:
	@go build -o build/todo cmd/todo/main.go

test:
	@go test -v ./...

run: build
	@./build/todo

migration:
	@migrate create -ext sql -dir migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
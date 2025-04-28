test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
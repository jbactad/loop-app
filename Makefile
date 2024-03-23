graphql-gen:
	@echo "Generating GraphQL code..."
	@gqlgen generate
	@echo "Done!"

generate:
	@echo "Generating code..."
	@go generate ./...
	@echo "Done!"

test-unit: graphql-gen generate
	@echo "Running unit tests..."
	@go test -v -cover ./...
	@echo "Done!"

run:
	@echo "Running dev server..."
	@go run main.go graphql
	@echo "Done!"

dep-up:
	@echo "Starting dependencies..."
	@docker-compose up -d
	@echo "Done!"

dep-down:
	@echo "Stopping dependencies..."
	@docker-compose down -v 
	@echo "Done!"

db-migrate:
	@echo "Running migrations..."
	@go run main.go db migrate
	@echo "Done!"
	
db-rollback:
	@echo "Rolling back migrations..."
	@go run main.go db rollback
	@echo "Done!"

db-seed:
	@echo
	@echo "Seeding database..."
	@go run main.go db seed
	@echo "Done!"

setup:
	@echo "Installing tools to $$(go env GOPATH)/bin"
	@echo "Make sure to add it in your PATH env var"
	@go install github.com/vektra/mockery/v2@v2.42.1
	@go install github.com/99designs/gqlgen@v0.17.45
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
	@echo "Done!"

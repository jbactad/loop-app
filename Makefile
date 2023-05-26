graphql-gen:
	@echo "Generating GraphQL code..."
	@go run github.com/99designs/gqlgen generate
	@echo "Done!"

generate:
	@echo "Generating code..."
	@go generate ./...
	@echo "Done!"

test-unit: generate
	@echo "Running unit tests..."
	@go test -v -cover ./...
	@echo "Done!"

graphql-gen:
	@echo "Generating GraphQL code..."
	@./tools/gqlgen generate
	@echo "Done!"

generate:
	@echo "Generating code..."
	@go generate ./...
	@echo "Done!"

test-unit: generate
	@echo "Running unit tests..."
	@go test -v -cover ./...
	@echo "Done!"

dev-run:
	@echo "Running dev server..."
	@go run main.go graphql
	@echo "Done!"

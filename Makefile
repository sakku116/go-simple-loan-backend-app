install-tools:
	@echo "Installing tools..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install gotest.tools/gotestsum@latest
	@echo "Tools installed."

gen-swagger:
	@echo "Generating swagger..."
	@swag init
	@echo "Swagger generated."

gen-mocks:
	@echo "Generating mocks..."
	@mockery --dir ./repository --all --output ./mocks
	@echo "Mocks generated."

run-tests:
	@echo "Running tests..."
	@gotestsum --format testname
	@echo "Tests passed."
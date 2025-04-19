build:
	@echo "Building the project..."
	@go build -o bin/main ./cmd/main.go

run:
	@echo "Running the project..."
	@go run ./cmd/main.go

clean:
	@echo "Cleaning up..."
	@rm -rf bin	

start:
	@echo "Starting the project from the binary..."
	@./bin/main
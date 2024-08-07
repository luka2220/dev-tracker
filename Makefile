# Format code 
fmt:
	go fmt ./...

# View possible issues in codebase
vet:
	go vet ./...

# Add any missing libraries and remove unsed ones
tidy: fmt
	go mod tidy

# Build the executable binary for the application
build:
	@go build -o bin/

# Run the root command 
run: build
	@./bin/devtasks

# Clean project files and remove current binary in ./bin
clean:
	go clean
	rm ./bin/devtasks

# Run the initalization tui test cases
test-init:
	go test ./tui/initialization -v

# Run the database test cases
test-db:
	go test ./database

# Display the database log file
log-db:
	cat ./database/db.log

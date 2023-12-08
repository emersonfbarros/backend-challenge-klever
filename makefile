BINARY_NAME=gowallet
TEST?=./...

# run app in debug mode
run:
	@go run ./main.go

# build app
build:
	@go build -o $(BINARY_NAME) main.go

# run app in release mode
release: build
	@GIN_MODE=release ./$(BINARY_NAME)

# tests
test:
	@go test $(TEST)

# verbose tests
test-verbose:
	@go test -v $(TEST)

clean:
	@rm -f $(BINARY_NAME)

# Variables
BINARY_NAME=bf
MAIN_PACKAGE=./main

all: clean build

clean:
	echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)

build: clean
	echo "Building..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)
	chmod +x $(BINARY_NAME)

run:
	echo "Running..."
	./$(BINARY_NAME)

#
# test:
#     echo "Running tests..."
#     go test ./...
#
# # Targets

.PHONY: clean build run

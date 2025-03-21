# Variables
BINARY_NAME=bf
MAIN_PACKAGE=./main
FILE?=./examples/hello.bf
FILE_NAME=$(notdir $(basename $(FILE) .bf))
ARGS?=""

default: clean build
	@echo "Running..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning..."
	(cd $(MAIN_PACKAGE) && go install)
	(cd $(MAIN_PACKAGE) && go clean)
	@rm -f $(BINARY_NAME)

build: clean
	@echo "Building..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)
	chmod +x $(BINARY_NAME)

run: build
	@echo "Running..."
	./$(BINARY_NAME) $(FILE) $(ARGS)

compile: build
	@echo "Compiling..."
	./$(BINARY_NAME) -c $(FILE)
	gcc $(FILE_NAME).c -o $(FILE_NAME)
	@rm $(FILE_NAME).c

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=myapp
LINTCMD=golangci-lint run

# Default target executed when no arguments are given to make
.PHONY: all
all: test build

# Build the project
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the application
.PHONY: run
run:
	$(GORUN) main.go

# Clean build files
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
.PHONY: test
test:
	$(GOTEST) -v ./...

# Lint the code
.PHONY: lint
lint:
	$(LINTCMD)

# Install dependencies
.PHONY: deps
deps:
	$(GOGET) -v ./...

# Format the code
.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

# Initialize the project (add more steps as needed)
.PHONY: init
init: deps

# Update dependencies
.PHONY: update
update:
	$(GOGET) -u ./...

# Show help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make              Build the project"
	@echo "  make build        Build the project"
	@echo "  make run          Run the application"
	@echo "  make clean        Clean build files"
	@echo "  make test         Run tests"
	@echo "  make lint         Lint the code"
	@echo "  make deps         Install dependencies"
	@echo "  make fmt          Format the code"
	@echo "  make init         Initialize the project"
	@echo "  make update       Update dependencies"
	@echo "  make help         Show this help message"

.PHONY: all
all: build run

.PHONY: run
run: build
	@echo "Running..."
	./bin/todo add

.PHONY: build
build:
	@echo "Building todo app..."
	go build -o bin/todo main.go
	@echo "Build completed!"
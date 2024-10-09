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

.PHONY: install
install:
	@echo "Installing todo_cli"
	go install
	@echo "todo_cli installed!"

.PHONY: uninstall
uninstall:
	@echo "Remove todo_cli"
	rm -f $(GOPATH)/bin/todo_cli
	@echo "todo_cli uninstalled!"
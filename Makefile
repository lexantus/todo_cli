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
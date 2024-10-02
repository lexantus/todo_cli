package main

import (
	"github.com/lexantus/todo_cli/cmd"
	"github.com/lexantus/todo_cli/tasks"
)

func main() {
	cmd.Execute()
	tasks.Init()
}

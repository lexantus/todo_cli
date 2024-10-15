package cmd

import (
	"fmt"
	"github.com/lexantus/todo_cli/storage"
	"github.com/lexantus/todo_cli/tasks"
	"github.com/spf13/cobra"
)

type Config struct {
	Tasks []tasks.Task `toml:"task"`
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List of todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var config Config
		storage.Read(&config)
		for _, task := range config.Tasks {
			fmt.Printf("ID: %s, Description: %s, Status: %d, Progress: %d%%\n", task.Id, task.Desc, task.Status, task.Progress)
		}
		for _, task := range config.Tasks {
			fmt.Println(task.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	flags := lsCmd.Flags()
	flags.BoolP("all", "a", false, "Print all tasks")
	flags.BoolP("completed", "c", false, "Print only completed tasks")
	flags.BoolP("removed", "r", false, "Print only removed tasks")
}

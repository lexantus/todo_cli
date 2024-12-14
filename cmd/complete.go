package cmd

import (
	"github.com/lexantus/todo_cli/storage"
	"github.com/lexantus/todo_cli/tasks"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete task with id",
	Run: func(cmd *cobra.Command, args []string) {
		storage := storage.NewStorage()
		storage.ChangeStatus(tasks.Id(args[0]), tasks.DONE)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

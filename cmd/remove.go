package cmd

import (
	"github.com/lexantus/todo_cli/storage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove task with ids",
	Run: func(cmd *cobra.Command, args []string) {
		storage := storage.NewStorage()
		storage.Delete(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

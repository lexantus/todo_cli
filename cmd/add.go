package cmd

import (
	"github.com/lexantus/todo_cli/tasks"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	PreRun: func(cmd *cobra.Command, args []string) {
		tasks.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args[:], " ")
		task := tasks.NewTask(description)
		err := tasks.Store(task)
		if err != nil {
			tasks.Logger.Error("Store from add command", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

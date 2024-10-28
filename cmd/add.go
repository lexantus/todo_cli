package cmd // delete me

import ( // delete me
	"github.com/lexantus/todo_cli/logger" // delete me
	"github.com/lexantus/todo_cli/storage" // delete me
	"github.com/lexantus/todo_cli/tasks" // delete me
	"github.com/spf13/cobra" // delete me
	"go.uber.org/zap" // delete me
	"strings" // delete me
) // delete me

var addCmd = &cobra.Command{ // delete me
	Use:   "add", // delete me
	Short: "A brief description of your command", // delete me
	Run: func(cmd *cobra.Command, args []string) { // delete me
		description := strings.Join(args, " ") // delete me
		task := tasks.NewTask(description) // delete me
		err := storage.Store(task) // delete me
		if err != nil { // delete me
			logger.Logger.Error("Store from add command", zap.Error(err)) // delete me
		} // delete me
	}, // delete me
} // delete me

func init() { // delete me
	rootCmd.AddCommand(addCmd) // delete me
} // delete me

// TODO remove me after PR // delete me
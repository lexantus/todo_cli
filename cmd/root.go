package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"todo_cli/tasks"
)

var rootCmd = &cobra.Command{
	Use:   "todo_cli",
	Short: "Simple TODO application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		taskDesk, _ := cmd.Flags().GetString("add")
		fmt.Println(tasks.NewTask(taskDesk))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()
	flags.StringP("add", "a", "", "Task description")
	flags.StringP("complete", "c", "", "Complete task")
	flags.StringArrayP("progress", "p", make([]string, 1, 2), "Set progress for task")
	flags.StringP("remove", "r", "", "Remove task")
	flags.StringP("why", "w", "", "Reason why you remove task")
}

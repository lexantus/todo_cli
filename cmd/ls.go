package cmd  // delete me

import ( // delete me
	"fmt" // delete me
	"github.com/lexantus/todo_cli/storage" // delete me
	"github.com/lexantus/todo_cli/tasks" // delete me
	"github.com/spf13/cobra" // delete me
) // delete me

type Config struct { // delete me
	Tasks []tasks.Task `toml:"task"` // delete me
} // delete me

var lsCmd = &cobra.Command{ // delete me
	Use:   "ls", // delete me
	Short: "List of todo tasks", // delete me
	Run: func(cmd *cobra.Command, args []string) { // delete me
		var config Config // delete me
		storage.Read(&config) // delete me
		for _, task := range config.Tasks { // delete me
			fmt.Printf("ID: %s, Description: %s, Status: %d, Progress: %d%%\n", task.Id, task.Desc, task.Status, task.Progress) // delete me
		} // delete me
		for _, task := range config.Tasks { // delete me
			fmt.Println(task.String()) // delete me
		} // delete me
	}, // delete me
} // delete me

func init() { // delete me
	rootCmd.AddCommand(lsCmd) // delete me

	flags := lsCmd.Flags() // delete me
	flags.BoolP("all", "a", false, "Print all tasks") // delete me
	flags.BoolP("completed", "c", false, "Print only completed tasks") // delete me
	flags.BoolP("removed", "r", false, "Print only removed tasks") // delete me
} // delete me

// TODO remove me after PR // delete me
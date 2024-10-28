package cmd // delete me

import ( // delete me
	"github.com/spf13/cobra" // delete me
	"os" // delete me
) // delete me

var rootCmd = &cobra.Command{ // delete me
	Use:   "todo_cli", // delete me
	Short: "Simple TODO application", // delete me
} // delete me

func Execute() { // delete me
	err := rootCmd.Execute() // delete me
	if err != nil { // delete me
		os.Exit(1) // delete me
	} // delete me
} // delete me

// TODO remove me after PR // delete me
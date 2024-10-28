package cmd // delete me

import ( // delete me
	"fmt" // delete me
	"github.com/spf13/cobra" // delete me
) // delete me

var completeCmd = &cobra.Command{ // delete me
	Use:   "complete", // delete me
	Short: "A brief description of your command", // delete me
	Run: func(cmd *cobra.Command, args []string) { // delete me
		fmt.Println("complete called") // delete me
	}, // delete me
} // delete me

func init() { // delete me
	rootCmd.AddCommand(completeCmd) // delete me
} // delete me

// TODO remove me after PR // delete me
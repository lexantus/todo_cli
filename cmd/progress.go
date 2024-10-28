package cmd // delete me

import ( // delete me
	"fmt" // delete me
	"github.com/spf13/cobra" // delete me
) // delete me

var progressCmd = &cobra.Command{ // delete me
	Use:   "progress", // delete me
	Short: "A brief description of your command", // delete me
	Long:  ``, // delete me
	Run: func(cmd *cobra.Command, args []string) { // delete me
		fmt.Println("progress called") // delete me
	}, // delete me
} // delete me

func init() { // delete me
	rootCmd.AddCommand(progressCmd) // delete me
} // delete me

// TODO remove me after PR // delete me
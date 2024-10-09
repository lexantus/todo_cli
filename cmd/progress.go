package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("progress called")
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}

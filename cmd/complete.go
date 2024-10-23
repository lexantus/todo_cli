package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

// TODO remove me after PR

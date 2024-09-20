/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo_cli/tasks"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tasks.NewTask("List me"))
		fmt.Println(tasks.NewTask("List me").Progress(10))
		fmt.Println(tasks.NewTask("List me").Complete("id3"))
		fmt.Println(tasks.NewTask("List me").Remove("id4", "reason why fail to do"))
		fmt.Println(tasks.NewTask("List me"))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	flags := lsCmd.Flags()
	flags.BoolP("all", "a", false, "Print all tasks")
	flags.BoolP("completed", "c", false, "Print only completed tasks")
	flags.BoolP("removed", "r", false, "Print only removed tasks")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goPhercise/task/db"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Great job you stud you! Now go study.")
			return
		}

		fmt.Println("You have the following tasks:")

		//index based so we want this to start at 1 not zero
		for i, task := range tasks {
			fmt.Printf("%d. %s, Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}

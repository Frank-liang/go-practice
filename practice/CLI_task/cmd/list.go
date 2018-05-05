package cmd

import (
	"fmt"
	"os"

	"github.com/Frank-liang/go/practice/CLI_task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complate! Why not take a vacation")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d:  %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

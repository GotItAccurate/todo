/*
Copyright © 2025 NAME <aadhii.yz@gmail.com>
*/

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a TODO.",
	Long:  `Add a TODO to the existing list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Provide a Task.\n")
			return
		}

		task := args[0]

		var deadline time.Time
		if len(args) > 1 {
			parsedDeadline, err := time.Parse("2006-01-02", args[1])
			if err != nil {
				fmt.Printf("Invalid deadline format. Use YYYY-MM-DD.\n")
				return
			}

			deadline = parsedDeadline
		} else {
			deadline = time.Now().Add(24 * time.Hour)
		}

		todos := loadTodos()

		todos = append(
			todos,
			Todo{
				Task:      task,
				CreatedAt: time.Now(),
				Deadline:  deadline,
				Done:      false,
			})

		saveTodos(todos)

		fmt.Printf("Added: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

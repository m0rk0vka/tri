/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/m0rk0vka/tri/todo"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			fmt.Errorf("%v", err)
		}

		for _, arg := range args {
			item := todo.Item{Text: arg}
			item.SetPriority(priority)
			items = append(items, item)
		}
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

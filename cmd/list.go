package cmd

import (
	"fmt"

	"github.com/gerlachry/doit/todo"
	"github.com/spf13/cobra"
)

var priorityMap = map[int]string{
	1: "H",
	2: "M",
	3: "L",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current outstanding tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.List(db)
		if err != nil {
			panic(err)
		}

		for _, t := range tasks {
			fmt.Printf("%-3d(%s)[%-10s]%-50s\n", t.ID, priorityMap[t.Priority], t.Project.Name, t.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

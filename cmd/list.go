package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gerlachry/doit/todo"
	"github.com/olekukonko/tablewriter"
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
		prj, err := cmd.Flags().GetString("project")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tasks, err := todo.List(db, prj)
		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			fmt.Printf("no open tasks for project %s\n", prj)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"ID", "Priority", "Project", "Task"})
		for _, t := range tasks {
			table.Append([]string{strconv.Itoa(t.ID), priorityMap[t.Priority], t.Project.Name, t.Name})
			//fmt.Printf("%-3d(%s)[%-10s]%-50s\n", t.ID, priorityMap[t.Priority], t.Project.Name, t.Name)
		}
		table.Render()
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
	listCmd.Flags().StringP("project", "p", "", "Project name to filter tasks on.")
}

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.List(db)
		if err != nil {
			panic(err)
		}

		for _, t := range tasks {
			// TODO: print with fixed width/spacing
			fmt.Printf("%d - (%s) [%s] %s\n", t.ID, priorityMap[t.Priority], t.Project.Name, t.Name)
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

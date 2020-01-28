package cmd

import (
	"fmt"
	"os"

	doit "github.com/gerlachry/doit/pkg"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo task",
	Long: `Add a new todo task:

Provide the task name and the priority
Example:
doit add --name="take the trash out" --priority=3	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("error parsing the name parameter: ", err)
			os.Exit(1)
		}

		priority, err := cmd.Flags().GetInt("priority")
		if err != nil {
			fmt.Println("error parsing the priority parameter", err)
			os.Exit(1)
		}
		t := doit.Task{Name: name, Priority: priority, Completed: 0, DB: db}

		err = t.Insert()
		if err != nil {
			os.Exit(1)
		}
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
	addCmd.Flags().String("name", "", "The task to do")
	addCmd.Flags().Int("priority", 3, "The priorit for the task, 1=High, 2=Medium, 3=Low")
}

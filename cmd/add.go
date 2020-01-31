package cmd

import (
	"fmt"
	"os"

	"github.com/gerlachry/doit/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:       "add",
	ValidArgs: []string{"name"},
	Short:     "Add a new todo task",
	Long: `Add a new todo task:

Provide the task name and the priority
Example:
doit add "take the trash out" --priority=3 --project=house`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		priority, err := cmd.Flags().GetInt("priority")
		if err != nil {
			fmt.Println("error parsing the priority parameter", err)
			os.Exit(1)
		}

		prj, err := cmd.Flags().GetString("project")
		if err != nil {
			fmt.Println("error parsing the project name", err)
			os.Exit(1)
		}

		p := todo.Project{Name: prj}
		t := todo.Task{Name: name, Priority: priority, Completed: 0, Project: p}

		err = t.Insert(db)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println("added task")
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
	//addCmd.Flags().String("name", "", "The task to do")
	addCmd.Flags().Int("priority", 3, "The priorit for the task, 1=High, 2=Medium, 3=Low")
	addCmd.Flags().String("project", "", "Optional project to assign the task to")
}

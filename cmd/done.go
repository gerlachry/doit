package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gerlachry/doit/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Args:  cobra.ExactValidArgs(1),
	Short: "Mark a task as done using the ID",
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID required as the first parameter")
			fmt.Println(err)
			os.Exit(1)
		}

		err = todo.Done(i, db)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Marked task %d as completed\n", i)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//doneCmd.Flags().Int("id", 0, "The task ID")
}

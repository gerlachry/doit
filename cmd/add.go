/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"stash.merck.com/gerlachr/doit/pkg"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called hopefully the db is still open...")
		err := db.Ping()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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
		t := pkg.Task{Name: name, Priority: priority, Completed: 0, DB: db}

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
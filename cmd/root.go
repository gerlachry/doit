package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/gerlachry/doit/todo"
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var db *sql.DB
var initDatabase bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doit",
	Short: "A bar bones todo cli application",
	Long:  "A bar bones todo cli application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer closeDB()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error from execute", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initDB)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doit.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&initDatabase, "initDatabase", "i", false, "init the database")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetDefault("db", "~/doit.db")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		//fmt.Println("setting config file from flag")
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".doit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".doit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initDB() {
	d, err := sql.Open("sqlite3", viper.GetString("db"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = d.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if initDatabase {
		fmt.Println("Creating backend tables...")
		err = initSchema(d)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	db = d
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(todo.InitDB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("initialized the backend database")
	return nil
}

func closeDB() {
	db.Close()
}

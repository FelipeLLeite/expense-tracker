/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the expenses Database.",
	Long: `This command will initialize the database for the client expenses.

	If the client database does not exist, it create one. If it does exist, it connect to it and perferm actions.
	If no database needs to be created, this command will return basic information form the client database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args %v\n", args)
		fmt.Printf("amount of args %d\n", len(args))
		// Log a structured message
		logger.Info("User logged in", "userID", 123, "action", "login")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().Int("userID", -1, "The userID, which it will be use to initialize the Id to extrac all information about the user expenses.")
}

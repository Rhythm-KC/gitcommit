/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitcommit",
	Short: "A interactive commit tool for git",
	Long: `gitcommit is an cli application to have better flow of commits.
When trying to commit multiple files we need to keep track of where they are,
and if you want to create multiple commits do this mulitple times.
In gitcommit, you have an interactive cli tool where you can tag files and
each is associated with a seperate commit.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

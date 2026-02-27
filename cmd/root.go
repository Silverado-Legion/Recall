package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "recall",
	Short: "Fast command-line tool for recalling notes and tasks right from the Terminal. Built with Go.",
}

func Execute() error {
	return rootCmd.Execute()
}
package cmd

import (
	"Recall/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new note.",
	Run: func(cmd *cobra.Command, args []string) {
		Note := strings.Join(args, " ")
		Title, _ := cmd.Flags().GetString("title")

		if (Title == "") {
			Title = utils.AutoTitle(Note)
		}

		fmt.Printf("Title: %v\nNote: %v", Title, Note)
	},
}

func init() {
	addCmd.Flags().String("title", "", "Custom title for the note.")
	rootCmd.AddCommand(addCmd)
}
package cmd

import (
	"Recall/internal/db"
	"Recall/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [content] --title [optional]",
	Short: "Create a new note.",
	Run: func(cmd *cobra.Command, args []string) {
		Note := strings.Join(args, " ")
		Title, _ := cmd.Flags().GetString("title")

		if (Note == "") {
			fmt.Println("Note cannot be empty.")
			return
		}

		if (Title == "") {
			Title = utils.AutoTitle(Note)
		}

		if err := db.AddNote(Title, Note); err != nil {
			fmt.Println("Error saving note:", err)
			return
		}
		fmt.Println("Note saved as:", Title)
	},
}

func init() {
	addCmd.Flags().String("title", "", "Custom title for the note.")
	rootCmd.AddCommand(addCmd)
}
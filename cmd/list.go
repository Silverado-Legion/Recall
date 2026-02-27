package cmd

import (
	"Recall/internal/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all note titles, their IDs, and their date and times.",
	Run: func(cmd *cobra.Command, args []string) {
		Notes, err := db.GetNotes()
		if err != nil {
			fmt.Println("Error fetching notes:", err)
			return
		}

		if len(Notes) == 0 {
			fmt.Println("Nothing to see here... just yet.\nTry creating a note with 'recall add Note content here'!")
		}

		fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		for _, note := range Notes {
			fmt.Printf("┃%3d. %-50s [%s]┃\n", note.ID, note.Title, note.CreatedAt)
		}
		fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
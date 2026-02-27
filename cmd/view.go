package cmd

import (
	"Recall/internal/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [id or title]",
	Short: "Finds a specific note (titles are case sensitive and must be exact)",
	Run: func(cmd *cobra.Command, args []string) {
		Target := strings.Join(args, " ")
		if id, err := strconv.Atoi(Target); err == nil {
			note, err := db.GetNote(id)
			if err != nil {
				fmt.Println("Error getting note:", err)
				return
			} 
			fmt.Printf("Title: %s [%d]\n...\n%s\n...\nCreated at: [%v]", note.Title, note.ID, note.Content, note.CreatedAt)
			return
		}

		note, duplicate, err := db.GetNode(Target)
			if (err != nil) {
				fmt.Println("Error getting note:", err)
				return
			}

			if (note == nil) {
				fmt.Printf("Couldn't find note: '%s'", Target)
				return
			}

			if (duplicate) {
				fmt.Println("Multiple notes with that title were found. The most recent note with that title will be displayed below. If you need a different note, view it by ID.\n...")
			}
			fmt.Printf("Title: %s [%d]\n...\n%s\n...\nCreated at: [%v]", note.Title, note.ID, note.Content, note.CreatedAt)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
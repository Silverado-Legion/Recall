package cmd

import (
	"Recall/internal/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [Note ID]",
	Short: "Toggles completion of a note based on ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if id, err := strconv.Atoi(args[0]); err == nil {
			note, err := db.GetNote(id)
			if err != nil {
				fmt.Printf("Error retrieving note: %s\n", err)
				return
			}

			err = db.CompleteNote(id, !note.Completed)
			if err != nil {
				fmt.Printf("Error marking note: %s\n", err)
				return
			}

			var status string
			if note.Completed {
				status = "INCOMPLETE"
			} else {
				status = "COMPLETE"
			}

			fmt.Printf("Marked note '%s' [%d]: %s\n", note.Title, note.ID, status)
		} else {
			fmt.Printf("Usage:\n%s\n%s", cmd.Use, cmd.Short)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

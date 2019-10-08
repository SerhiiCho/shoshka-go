package cmd

import (
	"github.com/SerhiiCho/shoshka-go/telegram"

	"github.com/spf13/cobra"
)

var titlesCmd = &cobra.Command{
	Use:   "titles",
	Short: "Checks new photo reports at specific target site by parsing html and checking if there are new posts added. If new added, it takes the title, link and image url and sends it to a telegram chat.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, msg := range telegram.GetMessagesWithNewReports() {
			telegram.SendMessage(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(titlesCmd)
}

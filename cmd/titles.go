package cmd

import (
	"github.com/SerhiiCho/shoshka-go/telegram"

	"github.com/spf13/cobra"
)

var titlesCmd = &cobra.Command{
	Use:   "titles",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, msg := range telegram.GetMessagesWithNewReports() {
			telegram.SendMessage(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(titlesCmd)
}

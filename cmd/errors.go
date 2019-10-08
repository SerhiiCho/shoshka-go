package cmd

import (
	"github.com/SerhiiCho/shoshka-go/telegram"

	"github.com/spf13/cobra"
)

var errorsCmd = &cobra.Command{
	Use:   "errors",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, msg := range telegram.GetMessagesWithNewErrors() {
			telegram.SendMessage(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(errorsCmd)
}

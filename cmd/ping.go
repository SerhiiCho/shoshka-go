package cmd

import (
	"github.com/SerhiiCho/shoshka-go/telegram"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, msg := range telegram.GetMessageIfPingIsNotSuccessful() {
			telegram.SendMessage(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

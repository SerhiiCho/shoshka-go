package cmd

import (
	"github.com/SerhiiCho/shoshka-go/telegram"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Checks the state of the production server by using ping command. Sends 3 ping requests and if 1 of 3 requests is not successful, sends telegram message to a provided chat.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, msg := range telegram.GetMessageIfPingIsNotSuccessful() {
			telegram.SendMessage(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

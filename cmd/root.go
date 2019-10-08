package cmd

import (
	"fmt"
	"github.com/SerhiiCho/shoshka-go/utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "shoshka-go",
	Short: "Shoshka is a telegram bot that checks for things like server log files, the state of the server and couple other things.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		utils.HandleError(err, "Error with cmd.Usage command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

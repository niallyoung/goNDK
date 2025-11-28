package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "giftwrap",
	Short: "Send and receive encrypted NOSTR messages",
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(receiveCmd)
	rootCmd.AddCommand(identityCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(queryCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

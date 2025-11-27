package cmd

import (
	"fmt"
	"os"

	"github.com/niallyoung/goNDK/examples/client-giftwrap/internal"
	"github.com/spf13/cobra"
)

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Show identity information",
	RunE:  runIdentity,
}

func init() {
	identityCmd.Flags().StringVarP(&keyPath, "key", "k", os.Getenv("HOME")+"/.nostr/key", "Path to private key")
}

func runIdentity(cmd *cobra.Command, args []string) error {
	id, err := internal.LoadOrGenerateKey(keyPath)
	if err != nil {
		return fmt.Errorf("load key: %w", err)
	}

	fmt.Printf("npub: %s\n", id.NPub)
	fmt.Printf("nsec: %s\n", id.Nsec)
	fmt.Printf("hex:  %s\n", id.PubKeyHex)
	
	return nil
}

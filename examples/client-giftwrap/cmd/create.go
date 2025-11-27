package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/niallyoung/goNDK/event"
	"github.com/niallyoung/goNDK/examples/client-giftwrap/internal"
	"github.com/niallyoung/goNDK/identity"
	"github.com/niallyoung/goNDK/nips/nip59"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create encrypted event without publishing (for testing)",
	RunE:  runCreate,
}

var (
	senderKey string
)

func init() {
	createCmd.Flags().StringVarP(&destination, "destination", "d", "", "Recipient npub (required)")
	createCmd.Flags().StringVarP(&message, "message", "m", "", "Text message")
	createCmd.Flags().StringVarP(&senderKey, "sender-key", "s", "", "Sender key path (optional, generates if not provided)")
	createCmd.MarkFlagRequired("destination")
	createCmd.MarkFlagRequired("message")
}

func runCreate(cmd *cobra.Command, args []string) error {
	var senderID *identity.ExtendedIdentity
	var err error

	if senderKey != "" {
		senderID, err = internal.LoadOrGenerateKey(senderKey)
		if err != nil {
			return fmt.Errorf("load sender key: %w", err)
		}
	} else {
		senderID, err = identity.Generate()
		if err != nil {
			return fmt.Errorf("generate sender key: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Generated sender: %s\n", senderID.NPub)
	}

	recipientPubKey, err := identity.NpubToHex(destination)
	if err != nil {
		return fmt.Errorf("invalid destination: %w", err)
	}

	msg := internal.EncodeTextMessage(message)
	content, err := internal.EncodeMessageJSON(msg)
	if err != nil {
		return fmt.Errorf("encode message: %w", err)
	}

	innerEvent := event.NewEvent(1, content, nil, nil, nil, &senderID.PubKeyHex, nil)
	if err := innerEvent.Sign(senderID.PrivKeyHex); err != nil {
		return fmt.Errorf("sign inner event: %w", err)
	}

	wrapped, err := nip59.Wrap(innerEvent, senderID.PrivKeyHex, recipientPubKey)
	if err != nil {
		return fmt.Errorf("wrap message: %w", err)
	}

	eventJSON, err := json.MarshalIndent(wrapped, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	fmt.Println(string(eventJSON))
	return nil
}

package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
	"github.com/niallyoung/goNDK/examples/client-giftwrap/internal"
	"github.com/niallyoung/goNDK/identity"
	"github.com/niallyoung/goNDK/nips/nip59"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send encrypted message or file",
	RunE:  runSend,
}

var (
	destination string
	message     string
	filename    string
	relay       string
	keyPath     string
)

func init() {
	sendCmd.Flags().StringVarP(&destination, "destination", "d", "", "Recipient npub (required)")
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "Text message to send")
	sendCmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename for piped data")
	sendCmd.Flags().StringVarP(&relay, "relay", "r", "wss://relay.damus.io", "Relay URL")
	sendCmd.Flags().StringVarP(&keyPath, "key", "k", os.Getenv("HOME")+"/.nostr/key", "Path to private key")
	sendCmd.MarkFlagRequired("destination")
}

func runSend(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	senderID, err := internal.LoadOrGenerateKey(keyPath)
	if err != nil {
		return fmt.Errorf("load key: %w", err)
	}

	recipientPubKey, err := identity.NpubToHex(destination)
	if err != nil {
		return fmt.Errorf("invalid destination: %w", err)
	}

	var msg internal.Message
	if message != "" {
		msg = internal.EncodeTextMessage(message)
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		if filename == "" {
			filename = "file.bin"
		}
		msg = internal.EncodeFileMessage(data, filename, "application/octet-stream")
	}

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

	rm := client.NewRelayManager(relay)
	if err := rm.Connect(ctx); err != nil {
		return fmt.Errorf("connect to relay: %w", err)
	}
	defer rm.Close()

	if _, err := rm.Publish(ctx, wrapped); err != nil {
		return fmt.Errorf("publish event: %w", err)
	}

	fmt.Printf("✅ Sent to %s\n", destination)
	fmt.Printf("Event ID: %s\n", *wrapped.ID)
	
	return nil
}

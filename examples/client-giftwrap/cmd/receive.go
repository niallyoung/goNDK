package cmd

import (
	"context"
	"encoding/json"
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

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive and decrypt message",
	RunE:  runReceive,
}

var (
	eventID string
	watch   bool
	output  string
)

func init() {
	receiveCmd.Flags().StringVarP(&eventID, "event-id", "e", "", "Event ID to fetch")
	receiveCmd.Flags().StringVarP(&relay, "relay", "r", "wss://relay.damus.io", "Relay URL")
	receiveCmd.Flags().StringVarP(&keyPath, "key", "k", os.Getenv("HOME")+"/.nostr/key", "Path to private key")
	receiveCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch for incoming messages")
	receiveCmd.Flags().StringVarP(&output, "output", "o", "", "Output directory for files")
}

func runReceive(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	recipientID, err := internal.LoadOrGenerateKey(keyPath)
	if err != nil {
		return fmt.Errorf("load key: %w", err)
	}

	if eventID != "" {
		return fetchAndDecrypt(ctx, recipientID.PrivKeyHex, eventID)
	}

	if watch {
		return watchMessages(ctx, recipientID)
	}

	return readFromStdin(recipientID.PrivKeyHex)
}

func fetchAndDecrypt(ctx context.Context, privKey, eventID string) error {
	fmt.Fprintf(os.Stderr, "🔌 Connecting to %s...\n", relay)
	rm := client.NewRelayManager(relay)
	if err := rm.Connect(ctx); err != nil {
		return fmt.Errorf("connect to relay: %w", err)
	}
	defer rm.Close()

	fmt.Fprintf(os.Stderr, "🔍 Fetching event %s...\n", eventID)
	ev, err := client.FetchEventByID(ctx, rm, eventID)
	if err != nil {
		return fmt.Errorf("fetch event: %w", err)
	}

	return decryptAndDisplay(ev, privKey)
}

func watchMessages(ctx context.Context, recipientID *identity.ExtendedIdentity) error {
	fmt.Fprintf(os.Stderr, "🔌 Connecting to %s...\n", relay)
	rm := client.NewRelayManager(relay)
	if err := rm.Connect(ctx); err != nil {
		return fmt.Errorf("connect to relay: %w", err)
	}
	defer rm.Close()
	fmt.Fprintf(os.Stderr, "✓ Connected\n")

	filter := client.Filter{
		Kinds: []int{1059},
	}

	sub, err := rm.Subscribe(ctx, []client.Filter{filter})
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	fmt.Println("👀 Watching for messages...")

	return sub.Receive(ctx, func(ctx context.Context, ev *event.Event) {
		if err := decryptAndDisplay(ev, recipientID.PrivKeyHex); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	})
}

func readFromStdin(privKey string) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read stdin: %w", err)
	}

	var ev event.Event
	if err := json.Unmarshal(data, &ev); err != nil {
		return fmt.Errorf("parse event: %w", err)
	}

	return decryptAndDisplay(&ev, privKey)
}

func decryptAndDisplay(ev *event.Event, privKey string) error {
	innerEvent, err := nip59.Unwrap(ev, privKey, "")
	if err != nil {
		return fmt.Errorf("unwrap: %w", err)
	}

	msg, err := internal.DecodeMessageJSON(innerEvent.Content)
	if err != nil {
		return fmt.Errorf("decode message: %w", err)
	}

	msgType, data, filename, err := internal.DecodeMessage(msg)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	if msgType == "text" {
		fmt.Printf("📨 Message: %s\n", string(data))
		return nil
	}

	if output != "" {
		path := output + "/" + filename
		if err := os.WriteFile(path, data, 0644); err != nil {
			return fmt.Errorf("write file: %w", err)
		}
		fmt.Printf("📎 File saved: %s (%d bytes)\n", path, len(data))
	} else {
		fmt.Printf("📎 File: %s (%d bytes)\n", filename, len(data))
		os.Stdout.Write(data)
	}

	return nil
}

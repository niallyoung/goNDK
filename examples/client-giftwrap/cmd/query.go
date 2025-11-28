package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query relay for gift wrap events",
	RunE:  runQuery,
}

var (
	limit int
)

func init() {
	queryCmd.Flags().StringVarP(&relay, "relay", "r", "wss://relay.damus.io", "Relay URL")
	queryCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Max events to fetch")
}

func runQuery(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Fprintf(os.Stderr, "🔌 Connecting to %s...\n", relay)
	rm := client.NewRelayManager(relay)
	if err := rm.Connect(ctx); err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer rm.Close()
	fmt.Fprintf(os.Stderr, "✓ Connected\n\n")

	filter := client.Filter{
		Kinds: []int{1059},
		Limit: limit,
	}

	sub, err := rm.Subscribe(ctx, []client.Filter{filter})
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	fmt.Fprintf(os.Stderr, "🔍 Querying for kind 1059 (Gift Wrap) events...\n\n")

	count := 0
	timeout := time.After(10 * time.Second)

	err = sub.Receive(ctx, func(ctx context.Context, ev *event.Event) {
		count++
		fmt.Printf("Event %d:\n", count)
		fmt.Printf("  ID: %s\n", *ev.ID)
		fmt.Printf("  Pubkey: %s\n", *ev.Pubkey)
		fmt.Printf("  Created: %s\n", time.Unix(int64(ev.CreatedAt), 0).Format(time.RFC3339))
		fmt.Printf("  Content: %d bytes\n", len(ev.Content))
		
		for _, tag := range ev.Tags {
			if len(tag) > 1 && tag[0] == "p" {
				fmt.Printf("  Recipient: %s\n", tag[1])
			}
		}
		fmt.Println()
	})

	select {
	case <-timeout:
		fmt.Fprintf(os.Stderr, "\n✅ Found %d gift wrap events\n", count)
		return nil
	case <-ctx.Done():
		fmt.Fprintf(os.Stderr, "\n✅ Found %d gift wrap events\n", count)
		return nil
	default:
		if err != nil {
			return err
		}
	}

	return nil
}

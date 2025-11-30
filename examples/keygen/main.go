// Package main provides a command-line tool for generating NOSTR keypairs.
//
// Usage:
//
//	keygen                    # Generate new keypair
//	keygen -format json       # Output as JSON
//	keygen -format bech32     # Output as nsec|npub (default)
//	keygen -format hex        # Output as hex private|public
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/niallyoung/goNDK/identity"
)

const (
	formatBech32 = "bech32"
	formatHex    = "hex"
	formatJSON   = "json"
)

type output struct {
	Nsec       string `json:"nsec"`
	NPub       string `json:"npub"`
	PrivKeyHex string `json:"privkey_hex"`
	PubKeyHex  string `json:"pubkey_hex"`
}

func main() {
	format := flag.String("format", formatBech32, "Output format: bech32, hex, or json")
	flag.Parse()

	// Generate keypair
	id, err := identity.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating keypair: %v\n", err)
		os.Exit(1)
	}

	// Output in requested format
	switch *format {
	case formatBech32:
		fmt.Printf("%s|%s\n", id.Nsec, id.NPub)
	case formatHex:
		fmt.Printf("%s|%s\n", id.PrivKeyHex, id.PubKeyHex)
	case formatJSON:
		out := output{
			Nsec:       id.Nsec,
			NPub:       id.NPub,
			PrivKeyHex: id.PrivKeyHex,
			PubKeyHex:  id.PubKeyHex,
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(out); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s\n", *format)
		fmt.Fprintf(os.Stderr, "Valid formats: bech32, hex, json\n")
		os.Exit(1)
	}
}

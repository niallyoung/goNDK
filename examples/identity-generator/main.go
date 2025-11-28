package main

import (
	"fmt"
	"os"

	"github.com/niallyoung/goNDK/identity"
)

func main() {
	id, err := identity.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("npub: %s\n", id.NPub)
	fmt.Printf("nsec: %s\n", id.Nsec)
	fmt.Printf("hex:  %s\n", id.PubKeyHex)
}

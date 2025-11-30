package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
	"testing"
)

func TestKeygenBech32(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("keygen failed: %v", err)
	}

	result := strings.TrimSpace(string(out))
	parts := strings.Split(result, "|")
	
	if len(parts) != 2 {
		t.Fatalf("expected nsec|npub format, got: %s", result)
	}

	nsec, npub := parts[0], parts[1]

	if !strings.HasPrefix(nsec, "nsec1") {
		t.Errorf("nsec should start with 'nsec1', got: %s", nsec)
	}

	if !strings.HasPrefix(npub, "npub1") {
		t.Errorf("npub should start with 'npub1', got: %s", npub)
	}

	if len(nsec) != 63 {
		t.Errorf("nsec should be 63 chars, got: %d", len(nsec))
	}

	if len(npub) != 63 {
		t.Errorf("npub should be 63 chars, got: %d", len(npub))
	}
}

func TestKeygenHex(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-format", "hex")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("keygen failed: %v", err)
	}

	result := strings.TrimSpace(string(out))
	parts := strings.Split(result, "|")
	
	if len(parts) != 2 {
		t.Fatalf("expected privkey|pubkey format, got: %s", result)
	}

	privKey, pubKey := parts[0], parts[1]

	if len(privKey) != 64 {
		t.Errorf("privkey should be 64 hex chars, got: %d", len(privKey))
	}

	if len(pubKey) != 64 {
		t.Errorf("pubkey should be 64 hex chars, got: %d", len(pubKey))
	}
}

func TestKeygenJSON(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-format", "json")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("keygen failed: %v", err)
	}

	var result output
	if err := json.Unmarshal(out, &result); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}

	if !strings.HasPrefix(result.Nsec, "nsec1") {
		t.Errorf("nsec should start with 'nsec1', got: %s", result.Nsec)
	}

	if !strings.HasPrefix(result.NPub, "npub1") {
		t.Errorf("npub should start with 'npub1', got: %s", result.NPub)
	}

	if len(result.PrivKeyHex) != 64 {
		t.Errorf("privkey_hex should be 64 chars, got: %d", len(result.PrivKeyHex))
	}

	if len(result.PubKeyHex) != 64 {
		t.Errorf("pubkey_hex should be 64 chars, got: %d", len(result.PubKeyHex))
	}
}

func TestKeygenInvalidFormat(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-format", "invalid")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected error for invalid format")
	}

	if !strings.Contains(stderr.String(), "Unknown format") {
		t.Errorf("expected 'Unknown format' error, got: %s", stderr.String())
	}
}

func TestKeygenUniqueness(t *testing.T) {
	// Generate two keys and ensure they're different
	cmd1 := exec.Command("go", "run", "main.go")
	out1, err := cmd1.Output()
	if err != nil {
		t.Fatalf("keygen 1 failed: %v", err)
	}

	cmd2 := exec.Command("go", "run", "main.go")
	out2, err := cmd2.Output()
	if err != nil {
		t.Fatalf("keygen 2 failed: %v", err)
	}

	if bytes.Equal(out1, out2) {
		t.Error("generated keys should be unique")
	}
}

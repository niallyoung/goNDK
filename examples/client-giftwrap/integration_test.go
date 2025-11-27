package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCLI_Help(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	cmd := exec.Command("./giftwrap", "--help")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err)
	
	assert.Contains(t, string(output), "Send and receive encrypted NOSTR messages")
	assert.Contains(t, string(output), "send")
	assert.Contains(t, string(output), "receive")
}

func TestCLI_SendHelp(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	cmd := exec.Command("./giftwrap", "send", "--help")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err)
	
	assert.Contains(t, string(output), "destination")
	assert.Contains(t, string(output), "message")
	assert.Contains(t, string(output), "relay")
}

func TestCLI_ReceiveHelp(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	cmd := exec.Command("./giftwrap", "receive", "--help")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err)
	
	assert.Contains(t, string(output), "event-id")
	assert.Contains(t, string(output), "watch")
	assert.Contains(t, string(output), "relay")
}

func TestCLI_KeyGeneration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "test.key")

	cmd := exec.Command("./giftwrap", "send", 
		"--destination", "npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6",
		"--message", "test",
		"--key", keyPath)
	
	// Will fail to connect but should generate key
	cmd.Run()
	
	assert.FileExists(t, keyPath)
	
	data, err := os.ReadFile(keyPath)
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(data), "nsec1"))
}

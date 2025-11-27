package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadOrGenerateKey(t *testing.T) {
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "test.key")
	
	id1, err := LoadOrGenerateKey(keyPath)
	require.NoError(t, err)
	require.NotNil(t, id1)
	
	assert.FileExists(t, keyPath)
	
	id2, err := LoadOrGenerateKey(keyPath)
	require.NoError(t, err)
	require.NotNil(t, id2)
	
	assert.Equal(t, id1.PrivKeyHex, id2.PrivKeyHex)
	assert.Equal(t, id1.PubKeyHex, id2.PubKeyHex)
}

func TestLoadKey(t *testing.T) {
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "test.key")
	
	id, err := LoadKey(keyPath)
	assert.Error(t, err)
	assert.Nil(t, id)
	
	err = os.WriteFile(keyPath, []byte("invalid"), 0600)
	require.NoError(t, err)
	
	id, err = LoadKey(keyPath)
	assert.Error(t, err)
	assert.Nil(t, id)
}

func TestSaveKey(t *testing.T) {
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "test.key")
	
	nsec := "nsec1vl029mgpspedva04g90vltkh6fvh240zqtv9k0t9af8935ke9laqsnlfe5"
	
	err := SaveKey(keyPath, nsec)
	require.NoError(t, err)
	
	data, err := os.ReadFile(keyPath)
	require.NoError(t, err)
	assert.Equal(t, nsec, string(data))
	
	info, err := os.Stat(keyPath)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
}

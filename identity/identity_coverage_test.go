package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIdentity(t *testing.T) {
	pubkey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	npub := "npub10xcvvlqfydr8zrmsj3hnswlsjtq92hcxg99vdtcxf4uc0cfe0s8s2whckx"
	
	id := NewIdentity(pubkey, npub)
	assert.Equal(t, pubkey, id.Pubkey)
	assert.Equal(t, npub, id.NPub)
}

func TestIdentityValidate(t *testing.T) {
	tests := []struct {
		name    string
		id      Identity
		wantErr bool
	}{
		{
			name: "valid identity",
			id: Identity{
				Pubkey: "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
				NPub:   "npub10xcvvlqfydr8zrmsj3hnswlsjtq92hcxg99vdtcxf4uc0cfe0s8s2whckx",
			},
			wantErr: false,
		},
		{
			name: "empty pubkey",
			id: Identity{
				Pubkey: "",
				NPub:   "npub10xcvvlqfydr8zrmsj3hnswlsjtq92hcxg99vdtcxf4uc0cfe0s8s2whckx",
			},
			wantErr: true,
		},
		{
			name: "invalid hex pubkey",
			id: Identity{
				Pubkey: "not-hex",
				NPub:   "npub10xcvvlqfydr8zrmsj3hnswlsjtq92hcxg99vdtcxf4uc0cfe0s8s2whckx",
			},
			wantErr: true,
		},
		{
			name: "wrong length pubkey",
			id: Identity{
				Pubkey: "79be667e",
				NPub:   "npub10xcvvlqfydr8zrmsj3hnswlsjtq92hcxg99vdtcxf4uc0cfe0s8s2whckx",
			},
			wantErr: true,
		},
		{
			name: "empty npub",
			id: Identity{
				Pubkey: "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
				NPub:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.id.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFromHexInvalidKey(t *testing.T) {
	_, err := FromHex("invalid")
	assert.Error(t, err)
}



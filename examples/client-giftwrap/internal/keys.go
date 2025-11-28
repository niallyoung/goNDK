package internal

import (
	"os"
	"strings"

	"github.com/niallyoung/goNDK/identity"
)

func LoadOrGenerateKey(path string) (*identity.ExtendedIdentity, error) {
	id, err := LoadKey(path)
	if err == nil {
		return id, nil
	}
	
	id, err = identity.Generate()
	if err != nil {
		return nil, err
	}
	
	if err := SaveKey(path, id.Nsec); err != nil {
		return nil, err
	}
	
	return id, nil
}

func LoadKey(path string) (*identity.ExtendedIdentity, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	nsec := strings.TrimSpace(string(data))
	return identity.FromNsec(nsec)
}

func SaveKey(path string, nsec string) error {
	return os.WriteFile(path, []byte(nsec), 0600)
}

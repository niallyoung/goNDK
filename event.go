package gondo

import (
	"cmp"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nbd-wtf/go-nostr"
)

// Eventer TODO what's a good name?
type Eventer interface {
	ID() string
	String() string
	Serialize() []byte
	Sign(privateKey string, signOpts ...schnorr.SignOption) error
	Validate() error
}

type Event struct {
	nostr.Event
	//nolint:all
	id        *string         `json:"id,omitempty"`
	CreatedAt nostr.Timestamp `json:"created_at,omitempty"`
	Kind      int             `json:"kind"`
	Tags      nostr.Tags      `json:"tags,omitempty"`
	Content   string          `json:"content"`
	Sig       string          `json:"sig"`
}

func NewEvent(pubkey string, createdAt nostr.Timestamp, kind int, tags nostr.Tags, content string, sig string) Eventer {
	event := Event{
		Event: nostr.Event{
			PubKey:    pubkey,
			CreatedAt: createdAt,
			Kind:      kind,
			Tags:      tags,
			Content:   content,
			Sig:       sig,
		},
	}

	_ = event.setId()

	return event
}

func (e Event) ID() string {
	return cmp.Or(*e.id, e.setId())
}

func (e Event) setId() string {
	p := &e
	h := sha256.Sum256(p.Serialize())
	*e.id = hex.EncodeToString(h[:])
	return *e.id
}

func (e Event) Serialize() []byte {
	p := &e
	return p.Serialize()
}

func (e Event) Sign(privateKey string, signOpts ...schnorr.SignOption) error {
	pointerE := &e
	return pointerE.Sign(privateKey, signOpts...)
}

func (e Event) Validate() error {
	if ok, _ := e.validateSig(); !ok {
		return errors.New("sig no valid")
	}

	return validation.ValidateStruct(&e,
		validation.Field(&e.id, validation.Required, validation.Length(1, 32)), // NFI yet, but // TODO consider accessor->generate->private attrib
		validation.Field(&e.PubKey, validation.Required, validation.Length(1, 64)),
		validation.Field(&e.CreatedAt, validation.Required),
		validation.Field(&e.Kind, validation.Required),
		validation.Field(&e.Tags, validation.Required),
		validation.Field(&e.Content, validation.Required),
		validation.Field(&e.Sig, validation.Required),
	)
}

func (e Event) validateSig() (bool, error) {
	return e.CheckSignature()
}

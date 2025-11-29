package event

import (
	"errors"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Eventer TODO what's a better name?
type Eventer interface {
	Serialize() []byte
	Sign(privateKey string, signOpts ...schnorr.SignOption) error
	String() string
	Validate() error
	ValidateComplete() error
	ValidateSignature() (bool, error)
}

type Event struct {
	Kind      int       `json:"kind"`
	Content   string    `json:"content"`
	Tags      Tags      `json:"tags"`
	CreatedAt Timestamp `json:"created_at"`
	ID        *string   `json:"id"`     // set by Sign()
	Pubkey    *string   `json:"pubkey"` // set by Sign()
	Sig       *string   `json:"sig"`    // set by Sign()
}

func NewEvent(kind int, content string, tags Tags, createdAt *int64, id *string, pubkey *string, sig *string) *Event {
	var timestamp Timestamp
	if createdAt == nil {
		timestamp = Now()
	} else {
		timestamp = Timestamp(*createdAt)
	}

	return &Event{
		Kind:      kind,
		Content:   content,
		Tags:      tags,
		CreatedAt: timestamp,
		ID:        id,
		Pubkey:    pubkey,
		Sig:       sig,
	}
}

func (e Event) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Kind, validation.Required),
		// Content can be empty for many event types (contact lists, reactions, etc.)
		validation.Field(&e.CreatedAt, validation.Required, validation.Min(0)),
		validation.Field(&e.ID, validation.When(&e.ID != nil, is.Hexadecimal, validation.Length(64, 64))),
		validation.Field(&e.Pubkey,
			validation.When(&e.Pubkey != nil, is.Hexadecimal, validation.Length(64, 64)),
		),
		validation.Field(&e.Sig,
			validation.When(&e.Sig != nil, is.Hexadecimal, validation.Length(128, 128)),
		),
	)
}

// ValidateComplete validates both structure and signature
func (e Event) ValidateComplete() error {
	if err := e.Validate(); err != nil {
		return err
	}

	if ok, err := e.ValidateSignature(); !ok {
		return errors.Join(err, errors.New("signature not valid"))
	}

	return nil
}

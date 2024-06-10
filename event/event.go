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

func (e *Event) Validate() error {
	if err := validation.ValidateStruct(&e,
		validation.Field(&e.Kind, validation.Required),
		validation.Field(&e.Content, validation.Required),
		validation.Field(&e.Tags, validation.When(e.Tags != nil, validation.Each(is.UTFLetterNumeric))),
		validation.Field(&e.CreatedAt, validation.Required, validation.Min(0)),                           // time.Time.Unix()
		validation.Field(&e.ID, validation.When(e.ID != nil, is.Hexadecimal, validation.Length(64, 64))), // hex, sha256(event.Serialize())
		validation.Field(&e.Pubkey, // hex, secp256k1 schnorr public key derived from Sign(privatekey, ...)
			validation.When(e.Pubkey != nil, is.Hexadecimal, validation.Length(64, 64)),
		),
		validation.Field(
			&e.Sig, // hex, pubkey signed serialization
			validation.When(e.Sig != nil, is.Hexadecimal, validation.Length(128, 128)),
		),
	); err != nil {
		return err
	}

	if ok, err := e.ValidateSignature(); !ok {
		return errors.Join(err, errors.New("signature not valid"))
	}

	return nil
}

package event

import (
	"cmp"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
)

// Eventer TODO what's a better name?
type Eventer interface {
	ID() string
	String() string
	Serialize() []byte
	Sign(privateKey string, signOpts ...schnorr.SignOption) error
	Validate() error
}

// Event is a fully-formed NOSTR Event, signed and ready for publishing
type Event struct {
	//nolint:all
	id        *string   `json:"id"`
	PubKey    string    `json:"pubkey"`
	CreatedAt time.Time `json:"created_at"`
	Kind      int       `json:"kind"`
	Tags      Tags      `json:"tags,omitempty"`
	Content   string    `json:"content"`
	Sig       string    `json:"sig"`
}

func NewEvent(id *string, pubkey string, createdAt int64, kind int, tags Tags, content string, sig string) (Eventer, error) {
	e := &Event{
		id:        id,
		PubKey:    pubkey,
		CreatedAt: time.Unix(createdAt, 0),
		Kind:      kind,
		Tags:      tags,
		Content:   content,
		Sig:       sig,
	}

	return e, nil
}

func (e Event) ID() string {
	return cmp.Or(*e.id, e.setID())
}

func (e Event) setID() string {
	*e.id = generateID(e)
	return *e.id
}

func generateID(e Event) string {
	checksum := sha256.Sum256(e.Serialize())
	return hex.EncodeToString(checksum[:])
}

func (e Event) Validate() error {
	if err := validation.ValidateStruct(&e,
		validation.Field(&e.id, validation.Required, is.Hexadecimal, validation.Length(64, 64)),     // hex, sha256 event serialization
		validation.Field(&e.PubKey, validation.Required, is.Hexadecimal, validation.Length(64, 64)), // hex, secp256k1 schnorr public key
		validation.Field(&e.CreatedAt, validation.Required, validation.Min(time.Unix(0, 0))),
		validation.Field(&e.Kind, validation.Required),                                                  // only ever positive: 0..N?
		validation.Field(&e.Tags, validation.When(e.Tags != nil, validation.Each(is.UTFLetterNumeric))), // symbols?
		validation.Field(&e.Content, validation.Required),                                               // always?
		validation.Field(&e.Sig, validation.Required, is.Hexadecimal, validation.Length(128, 128)),      // hex, pubkey signed serialization
	); err != nil {
		return err
	}

	// TODO think through the use-cases here, do we want a generic flexible Event for pre-signed, pre-id ...?
	// TODO or at this point enforce / construct only fully-formed and signed Events ...?
	//
	// TODO consider use-case specific validators, upon basic value Validate() minimum?: ValidateUnsigned()?

	if ok, err := e.ValidateSignature(); !ok {
		return errors.Wrap(err, "signature not valid")
	}

	if err := e.ValidateID(); err != nil {
		return errors.Wrap(err, "id not valid")
	}

	return nil
}

func (e Event) ValidateID() error {
	if e.id == nil {
		return errors.New("id is nil")
	}

	if *e.id != generateID(e) {
		return errors.New("id is invalid")
	}

	return nil
}

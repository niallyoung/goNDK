package event

import (
	"cmp"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mailru/easyjson"
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

type Event struct {
	id        *string   `json:"id,omitempty"` // TODO remove json and custom marshal/unmarshal?  no effect: //nolint:all
	PubKey    string    `json:"pubkey,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Kind      int       `json:"kind"`
	Tags      Tags      `json:"tags,omitempty"`
	Content   string    `json:"content,omitempty"`
	Sig       string    `json:"sig,omitempty"`
}

func NewEvent(pubkey string, createdAt time.Time, kind int, tags Tags, content string, sig string) Eventer {
	e := &Event{
		PubKey:    pubkey,
		CreatedAt: createdAt,
		Kind:      kind,
		Tags:      tags,
		Content:   content,
		Sig:       sig,
	}

	_ = e.setId()

	return e
}

func (e *Event) ID() string {
	return cmp.Or(*e.id, e.setId())
}

func (e *Event) setId() string {
	checksum := sha256.Sum256(e.Serialize())
	id := hex.EncodeToString(checksum[:])
	e.id = &id
	return id
}

func (e *Event) Validate() error {
	if err := validation.ValidateStruct(&e,
		validation.Field(&e.id, validation.Required, is.Hexadecimal, validation.Length(32, 32)),     // hex sha256
		validation.Field(&e.PubKey, validation.Required, is.Hexadecimal, validation.Length(64, 64)), // hex secp256k1
		validation.Field(&e.CreatedAt, validation.Required, validation.Min(time.Unix(0, 0)), validation.Max(time.Unix(math.MaxInt64, 0))),
		validation.Field(&e.Kind, validation.Required, is.Int),
		validation.Field(&e.Tags, validation.Required),
		validation.Field(&e.Content, validation.Required),
		validation.Field(&e.Sig, validation.Required, is.Hexadecimal),
	); err != nil {
		return err
	}

	if ok, err := e.validateSig(); !ok {
		return errors.Wrap(err, "sig not valid")
	}

	return nil
}

func (e *Event) validateSig() (bool, error) {
	return e.CheckSignature()
}

// Stringer interface, just returns the raw JSON as a string.
func (e *Event) String() string {
	j, _ := easyjson.Marshal(e)
	return string(j)
}

// Serialize outputs a byte array that can be hashed/signed to identify/authenticate.
// JSON encoding as defined in RFC4627.
func (e *Event) Serialize() []byte {
	// the serialization process is just putting everything into a JSON array
	// so the order is kept. See NIP-01
	dst := make([]byte, 0)

	// the header portion is easy to serialize
	// [0,"pubkey",created_at,kind,[
	dst = append(dst, []byte(
		fmt.Sprintf(
			"[0,\"%s\",%d,%d,",
			e.PubKey,
			e.CreatedAt.Unix(),
			e.Kind,
		))...)

	// tags
	dst = e.Tags.marshalTo(dst)
	dst = append(dst, ',')

	// content needs to be escaped in general as it is user generated.
	dst = escapeString(dst, e.Content)
	dst = append(dst, ']')

	return dst
}

// CheckSignature checks if the signature is valid for the id
// (which is a hash of the serialized event content).
// returns an error if the signature itself is invalid.
func (e *Event) CheckSignature() (bool, error) {
	// read and check pubkey
	pk, err := hex.DecodeString(e.PubKey)
	if err != nil {
		return false, fmt.Errorf("event pubkey '%s' is invalid hex: %w", e.PubKey, err)
	}

	pubkey, err := schnorr.ParsePubKey(pk)
	if err != nil {
		return false, fmt.Errorf("event has invalid pubkey '%s': %w", e.PubKey, err)
	}

	// read signature
	s, err := hex.DecodeString(e.Sig)
	if err != nil {
		return false, fmt.Errorf("signature '%s' is invalid hex: %w", e.Sig, err)
	}
	sig, err := schnorr.ParseSignature(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse signature: %w", err)
	}

	// check signature
	hash := sha256.Sum256(e.Serialize())
	return sig.Verify(hash[:], pubkey), nil
}

// Sign signs an event with a given privateKey.
func (e *Event) Sign(privateKey string, signOpts ...schnorr.SignOption) error {
	s, err := hex.DecodeString(privateKey)
	if err != nil {
		return fmt.Errorf("sign called with invalid private key '%s': %w", privateKey, err)
	}

	if e.Tags == nil {
		e.Tags = make(Tags, 0)
	}

	sk, pk := btcec.PrivKeyFromBytes(s)
	pkBytes := pk.SerializeCompressed()
	e.PubKey = hex.EncodeToString(pkBytes[1:])

	h := sha256.Sum256(e.Serialize())
	sig, err := schnorr.Sign(sk, h[:], signOpts...)
	if err != nil {
		return err
	}

	*e.id = hex.EncodeToString(h[:])
	e.Sig = hex.EncodeToString(sig.Serialize())

	return nil
}

// Escaping strings for JSON encoding according to RFC8259.
// Also encloses result in quotation marks "".
func escapeString(dst []byte, s string) []byte {
	dst = append(dst, '"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '"':
			// quotation mark
			dst = append(dst, []byte{'\\', '"'}...)
		case c == '\\':
			// reverse solidus
			dst = append(dst, []byte{'\\', '\\'}...)
		case c >= 0x20:
			// default, rest below are control chars
			dst = append(dst, c)
		case c == 0x08:
			dst = append(dst, []byte{'\\', 'b'}...)
		case c < 0x09:
			dst = append(dst, []byte{'\\', 'u', '0', '0', '0', '0' + c}...)
		case c == 0x09:
			dst = append(dst, []byte{'\\', 't'}...)
		case c == 0x0a:
			dst = append(dst, []byte{'\\', 'n'}...)
		case c == 0x0c:
			dst = append(dst, []byte{'\\', 'f'}...)
		case c == 0x0d:
			dst = append(dst, []byte{'\\', 'r'}...)
		case c < 0x10:
			dst = append(dst, []byte{'\\', 'u', '0', '0', '0', 0x57 + c}...)
		case c < 0x1a:
			dst = append(dst, []byte{'\\', 'u', '0', '0', '1', 0x20 + c}...)
		case c < 0x20:
			dst = append(dst, []byte{'\\', 'u', '0', '0', '1', 0x47 + c}...)
		}
	}
	dst = append(dst, '"')
	return dst
}

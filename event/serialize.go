package event

import (
	"fmt"

	"github.com/mailru/easyjson"
)

// Serialize outputs a byte array that can be hashed/signed to identify/authenticate.
// JSON encoding as defined in RFC4627.
func (e Event) Serialize() []byte {
	// the serialization process is just putting everything into a JSON array
	// so the order is kept. See NIP-01
	dst := make([]byte, 0)

	// the header portion is easy to serialize
	// [0,"pubkey",created_at,kind,[
	dst = append(dst, []byte(
		fmt.Sprintf(
			"[0,\"%s\",%d,%d,",
			*e.PubKey,
			e.CreatedAt,
			e.Kind,
		))...)

	// tags
	dst = e.Tags.marshalTo(dst)
	dst = append(dst, ',')

	// content needs to be escaped in general as it is user generated.
	dst = EscapeString(dst, e.Content)
	dst = append(dst, ']')

	return dst
}

// Escaping strings for JSON encoding according to RFC8259.
// Also encloses result in quotation marks "".
func EscapeString(dst []byte, s string) []byte {
	dst = append(dst, '"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '"': // quotation mark / double-quote
			dst = append(dst, []byte{'\\', '"'}...)
		case c == '\\': // reverse solidus / backslash
			dst = append(dst, []byte{'\\', '\\'}...)
		case c >= 0x20: // default (>= space)
			dst = append(dst, c)
		// control chars
		case c == 0x08: // backspace
			dst = append(dst, []byte{'\\', 'b'}...)
		case c < 0x09: // (< horizontal tabulation)
			dst = append(dst, []byte{'\\', 'u', '0', '0', '0', '0' + c}...)
		case c == 0x09: // horizontal tabulation
			dst = append(dst, []byte{'\\', 't'}...)
		case c == 0x0a: // line feed / newline
			dst = append(dst, []byte{'\\', 'n'}...)
		case c == 0x0c: // form feed
			dst = append(dst, []byte{'\\', 'f'}...)
		case c == 0x0d: // carriage return
			dst = append(dst, []byte{'\\', 'r'}...)
		case c < 0x10: // (< data link escape)
			dst = append(dst, []byte{'\\', 'u', '0', '0', '0', 0x57 + c}...)
		case c < 0x1a: // (< substitute)
			dst = append(dst, []byte{'\\', 'u', '0', '0', '1', 0x20 + c}...)
		case c < 0x20: // (< space)
			dst = append(dst, []byte{'\\', 'u', '0', '0', '1', 0x47 + c}...)
		}
	}
	dst = append(dst, '"')
	return dst
}

// String implements Stringer interface, returns raw JSON as a string.
func (e Event) String() string {
	j, _ := easyjson.Marshal(e)
	return string(j)
}

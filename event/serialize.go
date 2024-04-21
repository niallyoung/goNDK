package event

import (
	"fmt"
)

// Serialize outputs a []byte JSON array of the Event
func (e Event) Serialize() []byte {
	dst := make([]byte, 0)

	// header
	// [0,"pubkey",created_at,kind,[
	dst = append(dst, []byte(fmt.Sprintf("[0,%q,%d,%d,", *e.Pubkey, e.CreatedAt, e.Kind))...)

	// tags
	dst = e.Tags.marshalTo(dst)
	dst = append(dst, ',')

	// content
	dst = escapeString(dst, e.Content)
	dst = append(dst, ']')

	return dst
}

// escapeString for JSON encoding according to RFC8259
func escapeString(dst []byte, s string) []byte {
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

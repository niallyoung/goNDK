package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_escapeString(t *testing.T) {
	scenarios := []struct {
		Name          string
		givenDst      []byte
		givenContent  string
		expectedBytes []byte
	}{
		{
			"escaped quotations",
			[]byte("destination \"quoted\""),
			"content \"quoted\"",
			[]byte("destination \"quoted\"\"content \\\"quoted\\\"\""),
		},
		{
			"escaped back-slash",
			[]byte("destination \\back-slashed\\"),
			"content \\\\double-slashed\\\\",
			[]byte("destination \\back-slashed\\\"content \\\\\\\\double-slashed\\\\\\\\\""),
		},
		{
			"0x21",
			[]byte{},
			string([]byte{0x21}),
			[]byte{0x22, 0x21, 0x22},
		},
		{
			"0x08",
			[]byte{},
			string([]byte{0x08}),
			[]byte{0x22, 0x5c, 0x62, 0x22},
		},
		{
			"0x09",
			[]byte{},
			string([]byte{0x09}),
			[]byte{0x22, 0x5c, 0x74, 0x22},
		},
		{
			"0x0a",
			[]byte{},
			string([]byte{0x0a}),
			[]byte{0x22, 0x5c, 0x6e, 0x22},
		},
		{
			"0x0c",
			[]byte{},
			string([]byte{0x0c}),
			[]byte{0x22, 0x5c, 0x66, 0x22},
		},
		{
			"0x0d",
			[]byte{},
			string([]byte{0x0d}),
			[]byte{0x22, 0x5c, 0x72, 0x22},
		},
		{
			"0x10",
			[]byte{},
			string([]byte{0x10}),
			[]byte{0x22, 0x5c, 0x75, 0x30, 0x30, 0x31, 0x30, 0x22},
		},
		{
			"0x1a",
			[]byte{},
			string([]byte{0x1a}),
			[]byte{0x22, 0x5c, 0x75, 0x30, 0x30, 0x31, 0x61, 0x22},
		},
		{
			"0x20",
			[]byte{},
			string([]byte{0x20}),
			[]byte{0x22, 0x20, 0x22},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			assert.Equal(t, s.expectedBytes, EscapeString(s.givenDst, s.givenContent))
			assert.Equal(t, string(s.expectedBytes), string(EscapeString(s.givenDst, s.givenContent)))
		})
	}
}

package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag(t *testing.T) {
	t.Run("empty Tag marshalTo", func(t *testing.T) {
		tag := Tag{}
		var b []byte
		tag.marshalTo(b)
		assert.Equal(t, ([]byte(nil)), b)
	})

	// TBC
	//t.Run("non-empty Tag marshalTo", func(t *testing.T) {
	//	tag := Tag{"first", "second"}
	//	var b []byte
	//	tag.marshalTo(b)
	//	assert.Equal(t, []byte(nil), b)
	//})
}

func TestTags(t *testing.T) {
	t.Run("", func(t *testing.T) {
		tags := make(Tags, 0, 1)
		tags = append(tags, Tag{"first", "second"})
		assert.Equal(t, Tags{Tag{"first", "second"}}, tags)
	})
}

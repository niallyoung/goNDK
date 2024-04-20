package event

type Tags []Tag

// marshalTo appends the JSON encoded byte of Tags
func (tags Tags) marshalTo(dst []byte) []byte {
	dst = append(dst, '[')
	for i, tag := range tags {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = tag.marshalTo(dst)
	}
	dst = append(dst, ']')
	return dst
}

type Tag []string

// marshalTo appends the JSON encoded byte of Tag
func (tag Tag) marshalTo(dst []byte) []byte {
	dst = append(dst, '[')
	for i, s := range tag {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = escapeString(dst, s)
	}
	dst = append(dst, ']')
	return dst
}

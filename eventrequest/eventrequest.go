package eventrequest

import "github.com/niallyoung/goNDK/event"

// EventRequest is for handler receipt and JSON unmarshal, before constructing an event.Event{}
type EventRequest struct {
	ID        string     `json:"id"`
	PubKey    string     `json:"pubkey"`
	CreatedAt int64      `json:"created_at"`
	Kind      int        `json:"kind"`
	Tags      event.Tags `json:"tags"`
	Content   string     `json:"content"`
	Sig       string     `json:"sig"`
}

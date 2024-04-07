package eventrequest

import "github.com/niallyoung/goNDK/event"

// EventRequest is intended for initial receipt and JSON unmarshalling, before constructing an Event{}
type EventRequest struct {
	ID        string     `json:"id"`
	PubKey    string     `json:"pubkey"`
	CreatedAt int64      `json:"created_at"`
	Kind      int        `json:"kind"`
	Tags      event.Tags `json:"tags"`
	Content   string     `json:"content"`
	Sig       string     `json:"sig"`
}

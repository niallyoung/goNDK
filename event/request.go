package event

// EventRequest is intended for relay handlers and their concerns:
//   - receipt, initial validation, etc. before transformation to the full Event type
type EventRequest struct {
	ID        string `json:"id"`
	PubKey    string `json:"pubkey"`
	CreatedAt int64  `json:"created_at"`
	Kind      int    `json:"kind"`
	Tags      Tags   `json:"tags"`
	Content   string `json:"content"`
	Sig       string `json:"sig"`
}

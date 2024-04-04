package event

import "time"

// EventRequest is intended for relay handlers and their concerns:
//   - receipt, initial validation, etc. before transformation to the full Event type
type EventRequest struct {
	ID        string    `json:"id"`
	PubKey    string    `json:"pubkey"`
	CreatedAt Timestamp `json:"created_at"`
	Kind      int       `json:"kind"`
	Tags      Tags      `json:"tags"`
	Content   string    `json:"content"`
	Sig       string    `json:"sig"`
}

// TODO I'm still not sold on this, prefer time.Time and serialize, will circle back
type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

func Now() Timestamp {
	return Timestamp(time.Now().Unix())
}

package client

import (
	"encoding/json"
	"github.com/niallyoung/goNDK/event"
)

type MessageType string

const (
	MessageTypeEvent  MessageType = "EVENT"  // NIP-01
	MessageTypeReq    MessageType = "REQ"    // NIP-01
	MessageTypeClose  MessageType = "CLOSE"  // NIP-01
	MessageTypeNotice MessageType = "NOTICE" // NIP-01
	MessageTypeEOSE   MessageType = "EOSE"   // NIP-01
	MessageTypeOK     MessageType = "OK"     // NIP-20
)

// EventMessage is used to publish and receive events
type EventMessage struct {
	SubscriptionID string // optional
	Event          *event.Event
}

func (m *EventMessage) MarshalJSON() ([]byte, error) {
	body := []any{MessageTypeEvent}
	if subID := m.SubscriptionID; subID != "" {
		body = append(body, subID)
	}
	body = append(body, m.Event)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ReqMessage is a request message.
// It's used to request events and subscribe to new updates.
type ReqMessage struct {
	SubscriptionID string
	Filters        []Filter
}

func (m *ReqMessage) MarshalJSON() ([]byte, error) {
	body := []any{MessageTypeReq, m.SubscriptionID}
	for _, f := range m.Filters {
		body = append(body, f)
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// CloseMessage is a close message.
// It's used to stop previous subscriptions.
type CloseMessage struct {
	SubscriptionID string
}

func (m *CloseMessage) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal([]any{MessageTypeClose, m.SubscriptionID})
	if err != nil {
		return nil, err
	}
	return b, nil
}

// NoticeMessage is a notice message.
// It's used to send human-readable messages to clients.
type NoticeMessage struct {
	Message string
}

// EOSEMessage is a EOSE message.
// It's used to notify clients all stored events have been sent.
type EOSEMessage struct {
	SubscriptionID string
}

// OKMessage is an OK message.
// It's used to notify clients if an EVENT was successful.
type OKMessage struct {
	EventID string
	OK      bool
	Message string
}

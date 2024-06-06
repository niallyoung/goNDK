package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync"

	"nhooyr.io/websocket"

	"github.com/niallyoung/goNDK/event"
)

type RelayManager struct {
	URL        string
	conn       *websocket.Conn
	doneChan   chan struct{}
	noticeChan chan string
	subMap     sync.Map // map[string]*subChannelGroup
	eventMap   sync.Map // map[string]*eventChannelGroup
}

func NewRelayManager(url string) RelayManager {
	return RelayManager{
		URL:        url,
		conn:       nil,
		doneChan:   make(chan struct{}),
		noticeChan: make(chan string, 100),
	}
}

func (rm *RelayManager) Connect(ctx context.Context) error {
	conn, _, err := websocket.Dial(ctx, rm.URL, nil)
	if err != nil {
		return errors.Join(err, errors.New("websocket connection"))
	}

	conn.SetReadLimit(math.MaxInt64 - 1) // disable read limit

	rm.conn = conn

	go func() {
	outer:
		for {
			select {
			case <-rm.doneChan:
				break outer
			default:
			}

			if err := rm.ReadMessage(ctx); err != nil {
				continue
			}
		}
	}()

	return nil
}

func (rm *RelayManager) WriteMessage(ctx context.Context, message json.Marshaler) error {
	body, err := message.MarshalJSON()
	if err != nil {
		return err
	}

	err = rm.conn.Write(ctx, websocket.MessageText, body)
	if err != nil {
		return err
	}

	return nil
}

func (rm *RelayManager) ReadMessage(ctx context.Context) error {
	_, b, err := rm.conn.Read(ctx)
	if err != nil {
		return err
	}

	var message []json.RawMessage
	if err = json.Unmarshal(b, &message); err != nil {
		return err
	}
	if len(message) == 0 {
		return errors.New("empty message")
	}
	var typ string
	if err = json.Unmarshal(message[0], &typ); err != nil {
		return err
	}

	switch typ {

	case string(MessageTypeNotice):
		if len(message) != 2 {
			return fmt.Errorf("invalid notice message length: %d", len(message))
		}
		var m string
		if err = json.Unmarshal(message[1], &m); err != nil {
			return err
		}
		rm.handleNoticeMessage(&NoticeMessage{Message: m})
		return nil

	case string(MessageTypeEvent):
		if len(message) != 3 {
			return fmt.Errorf("invalid event message length: %d", len(message))
		}
		var subID string
		if err = json.Unmarshal(message[1], &subID); err != nil {
			return err
		}
		var event event.Event
		if err = json.Unmarshal(message[2], &event); err != nil {
			return err
		}
		rm.handleEventMessage(&EventMessage{
			SubscriptionID: subID,
			Event:          &event,
		})
		return nil

	case string(MessageTypeEOSE):
		if len(message) != 2 {
			return fmt.Errorf("invalid EOSE message length: %d", len(message))
		}
		var subID string
		if err = json.Unmarshal(message[1], &subID); err != nil {
			return err
		}
		rm.handleEOSEMessage(&EOSEMessage{SubscriptionID: subID})
		return nil

	case string(MessageTypeOK):
		if len(message) != 4 {
			return fmt.Errorf("invalid OK message length: %d", len(message))
		}
		var eventID string
		if err = json.Unmarshal(message[1], &eventID); err != nil {
			return err
		}
		var ok bool
		if err = json.Unmarshal(message[2], &ok); err != nil {
			return err
		}
		var m string
		if err = json.Unmarshal(message[3], &m); err != nil {
			return err
		}
		rm.handleOKMessage(&OKMessage{
			EventID: eventID,
			OK:      ok,
			Message: m,
		})
		return nil
	}

	return fmt.Errorf("unsupported message type: %s", typ)
}

func (rm *RelayManager) handleNoticeMessage(m *NoticeMessage) error {
	select {
	case rm.noticeChan <- m.Message:
	default:
		// drop message
	}
	return nil
}

func (rm *RelayManager) handleEventMessage(m *EventMessage) error {
	value, ok := rm.subMap.Load(m.SubscriptionID)
	if !ok {
		return fmt.Errorf("unaddressed event message: subscription id: %s", m.SubscriptionID)
	}
	group, ok := value.(*subChannelGroup)
	if !ok {
		return errors.New("invalid value in subsciption map")
	}

	select {
	case group.eventChan <- m.Event:
	default:
		// drop message
	}
	return nil
}

func (rm *RelayManager) handleEOSEMessage(m *EOSEMessage) error {
	value, ok := rm.subMap.Load(m.SubscriptionID)
	if !ok {
		return fmt.Errorf("unaddressed EOSE message: subscription id: %s", m.SubscriptionID)
	}
	group, ok := value.(*subChannelGroup)
	if !ok {
		return errors.New("invalid value in subsciption map")
	}

	select {
	case group.eoseChan <- struct{}{}:
	default:
		// drop message
	}
	return nil
}

func (rm *RelayManager) handleOKMessage(m *OKMessage) error {
	value, ok := rm.eventMap.Load(m.EventID)
	if !ok {
		return fmt.Errorf("unaddressed OK message: event id: %s", m.EventID)
	}
	group, ok := value.(*eventChannelGroup)
	if !ok {
		return errors.New("invalid value in event map")
	}

	select {
	case group.okChan <- &CommandResult{
		OK:      m.OK,
		Message: m.Message,
	}:
	default:
		// drop message
	}
	return nil
}

type subChannelGroup struct {
	eventChan chan<- *event.Event
	eoseChan  chan<- struct{}
}

type eventChannelGroup struct {
	okChan chan<- *CommandResult
}

type CommandResult struct {
	OK      bool
	Message string
}

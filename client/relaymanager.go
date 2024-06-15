package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/google/uuid"
	"nhooyr.io/websocket"

	"github.com/niallyoung/goNDK/event"
)

type RelayManager struct {
	URL        string
	conn       *websocket.Conn
	doneChan   chan struct{}
	noticeChan chan string
	subMap     sync.Map // map[string]*subscriptionChannelGroup // TODO shift off mutex
	eventMap   sync.Map // map[string]*eventChannelGroup        // TODO shift off mutex
	closeOnce  sync.Once
	closeErr   error
}

func NewRelayManager(url string) *RelayManager {
	return &RelayManager{
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

type CommandResult struct {
	OK      bool
	Message string
}

type eventChannelGroup struct {
	okChan chan<- *CommandResult
}

// Publish submits an event to the relay server and waits for the command result.
func (rm *RelayManager) Publish(ctx context.Context, event *event.Event) (*CommandResult, error) {
	id := event.ID
	okChan := make(chan *CommandResult, 1)

	rm.eventMap.Store(id, &eventChannelGroup{
		okChan: okChan,
	})
	defer rm.eventMap.Delete(id)

	if err := rm.WriteMessage(ctx, &EventMessage{Event: event}); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("missing command result: %w", ctx.Err())
	case result := <-okChan:
		return result, nil
	}
}

type subscriptionChannelGroup struct {
	eventChan chan<- *event.Event
	eoseChan  chan<- struct{}
}

// Subscribe creates a subscription to the relay server with the given filters.
func (rm *RelayManager) Subscribe(_ context.Context, filters []Filter) (*Subscription, error) {
	if len(filters) == 0 {
		return nil, errors.New("at least one filter is required")
	}

	id := uuid.New().String()
	eventChan := make(chan *event.Event)
	eoseChan := make(chan struct{}, 1)

	trigger := func(ctx context.Context) error {
		// register subscription to client
		rm.subMap.Store(id, &subscriptionChannelGroup{
			eventChan: eventChan,
			eoseChan:  eoseChan,
		})

		req := ReqMessage{
			SubscriptionID: id,
			Filters:        filters,
		}
		if err := rm.WriteMessage(ctx, &req); err != nil {
			// unregister subscription from client
			rm.subMap.Delete(id)
			return err
		}
		return nil
	}

	closer := func(ctx context.Context) error {
		req := CloseMessage{SubscriptionID: id}
		if err := rm.WriteMessage(ctx, &req); err != nil {
			return err
		}

		// unregister subscription from client
		rm.subMap.Delete(id)
		return nil
	}

	return &Subscription{
		id:        id,
		eventChan: eventChan,
		eoseChan:  eoseChan,
		trigger:   trigger,
		closer:    closer,
	}, nil
}

// Notice returns a channel that receives notice messages from the relay server.
func (rm *RelayManager) Notice() <-chan string {
	return rm.noticeChan
}

// Close closes the client connection.
func (rm *RelayManager) Close() error {
	rm.closeOnce.Do(func() {
		close(rm.doneChan)

		err := rm.conn.Close(websocket.StatusNormalClosure, "")
		if err != nil {
			rm.closeErr = err
			return
		}
	})
	return rm.closeErr
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
		if err = rm.handleNoticeMessage(&NoticeMessage{Message: m}); err != nil {
			return err
		}
		return nil

	case string(MessageTypeEvent):
		if len(message) != 3 {
			return fmt.Errorf("invalid event message length: %d", len(message))
		}
		var subID string
		if err = json.Unmarshal(message[1], &subID); err != nil {
			return err
		}
		var e event.Event
		if err = json.Unmarshal(message[2], &e); err != nil {
			return err
		}
		if err = rm.handleEventMessage(&EventMessage{
			SubscriptionID: subID,
			Event:          &e,
		}); err != nil {
			return err
		}
		return nil

	case string(MessageTypeEOSE):
		if len(message) != 2 {
			return fmt.Errorf("invalid EOSE message length: %d", len(message))
		}
		var subID string
		if err = json.Unmarshal(message[1], &subID); err != nil {
			return err
		}
		if err = rm.handleEOSEMessage(&EOSEMessage{SubscriptionID: subID}); err != nil {
			return err
		}
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
		if err = rm.handleOKMessage(&OKMessage{
			EventID: eventID,
			OK:      ok,
			Message: m,
		}); err != nil {
			return err
		}
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
	group, ok := value.(*subscriptionChannelGroup)
	if !ok {
		return errors.New("invalid value in subscription map")
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
	group, ok := value.(*subscriptionChannelGroup)
	if !ok {
		return errors.New("invalid value in subscription map")
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

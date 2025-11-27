package client

import (
	"context"
	"errors"
	"time"

	"github.com/niallyoung/goNDK/event"
)

var ErrEventNotFound = errors.New("event not found")

// FetchEventByID fetches a single event by ID from the relay
func FetchEventByID(ctx context.Context, rm *RelayManager, eventID string) (*event.Event, error) {
	filter := Filter{
		IDs:   []string{eventID},
		Limit: 1,
	}

	sub, err := rm.Subscribe(ctx, []Filter{filter})
	if err != nil {
		return nil, err
	}

	fetchCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resultChan := make(chan *event.Event, 1)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-fetchCtx.Done():
				errChan <- ErrEventNotFound
				return
			case e := <-sub.eventChan:
				if e.ID != nil && *e.ID == eventID {
					resultChan <- e
					return
				}
			case <-sub.eoseChan:
				errChan <- ErrEventNotFound
				return
			}
		}
	}()

	select {
	case e := <-resultChan:
		return e, nil
	case err := <-errChan:
		return nil, err
	}
}

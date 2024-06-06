package client_test

import (
	"testing"
	"time"

	"github.com/aws/smithy-go/ptr"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

func TestEventMessageMarshalJSON(t *testing.T) {
	e := &event.Event{
		Kind:      event.KindTextNote,
		Content:   "short text note",
		Tags:      event.Tags{},
		CreatedAt: event.Timestamp(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ID:        ptr.String("f926f58579b974014c091f4d945e8e3de7f3f87bbc4a0b6a49f2b3d68be2c89d"),
		Pubkey:    ptr.String("7e7e9c42a91bfef19fa929e5fda1b72e0ebc1a4c1141673e2794234d86addf4e"),
		Sig:       ptr.String("7903b45c7863f053bb1e84e6308c0de6f2dd212a9496b2391c83859fec17a3f28427ce74e59deef34ff5c418d871601eb4b8c7a81390f4a3ccb08ba4bce55710"),
	}

	t.Run("without subscription id", func(t *testing.T) {
		message := client.EventMessage{
			Event: e,
		}

		expected := `[` +
			`"EVENT",` +
			`{"id":"f926f58579b974014c091f4d945e8e3de7f3f87bbc4a0b6a49f2b3d68be2c89d","pubkey":"7e7e9c42a91bfef19fa929e5fda1b72e0ebc1a4c1141673e2794234d86addf4e","created_at":1672531200,"kind":1,"tags":[],"content":"short text note","sig":"7903b45c7863f053bb1e84e6308c0de6f2dd212a9496b2391c83859fec17a3f28427ce74e59deef34ff5c418d871601eb4b8c7a81390f4a3ccb08ba4bce55710"}` +
			`]`

		b, err := message.MarshalJSON()
		if err != nil {
			t.Fatalf("message.MarshalJSON() failed: %s", err)
		}
		if string(b) != expected {
			t.Errorf("message.MarshalJSON() failed: expected %s, got %s", expected, string(b))
		}
	})

	t.Run("with subscription id", func(t *testing.T) {
		message := client.EventMessage{
			SubscriptionID: "sub-id",
			Event:          e,
		}

		expected := `[` +
			`"EVENT",` +
			`"sub-id",` +
			`{"id":"f926f58579b974014c091f4d945e8e3de7f3f87bbc4a0b6a49f2b3d68be2c89d","pubkey":"7e7e9c42a91bfef19fa929e5fda1b72e0ebc1a4c1141673e2794234d86addf4e","created_at":1672531200,"kind":1,"tags":[],"content":"short text note","sig":"7903b45c7863f053bb1e84e6308c0de6f2dd212a9496b2391c83859fec17a3f28427ce74e59deef34ff5c418d871601eb4b8c7a81390f4a3ccb08ba4bce55710"}` +
			`]`

		b, err := message.MarshalJSON()
		if err != nil {
			t.Fatalf("message.MarshalJSON() failed: %s", err)
		}
		if string(b) != expected {
			t.Errorf("message.MarshalJSON() failed: expected %s, got %s", expected, string(b))
		}
	})
}

func TestReqMessageMarshalJSON(t *testing.T) {
	message := client.ReqMessage{
		SubscriptionID: "sub-id",
		Filters: []client.Filter{
			{
				Kinds:  []int{event.KindTextNote},
				Search: "target text",
			},
		},
	}

	expected := `["REQ","sub-id",{"kinds":[1],"search":"target text"}]`

	b, err := message.MarshalJSON()
	if err != nil {
		t.Fatalf("message.MarshalJSON() failed: %s", err)
	}
	if string(b) != expected {
		t.Errorf("message.MarshalJSON() failed: expected %s, got %s", expected, string(b))
	}
}

func TestCloseMessageMarshalJSON(t *testing.T) {
	message := client.CloseMessage{
		SubscriptionID: "sub-id",
	}

	expected := `["CLOSE","sub-id"]`

	b, err := message.MarshalJSON()
	if err != nil {
		t.Fatalf("message.MarshalJSON() failed: %s", err)
	}
	if string(b) != expected {
		t.Errorf("message.MarshalJSON() failed: expected %s, got %s", expected, string(b))
	}
}

package event_test

import (
	"github.com/aws/smithy-go/ptr"

	"github.com/niallyoung/goNDK/event"
)

var ValidEvent = func() *event.Event {
	return event.NewEvent(
		1,
		"GM nostr welcome to Saturday!",
		event.Tags(nil),
		ptr.Int64(1712350548),
		ptr.String("b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76"),
		ptr.String("234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983"),
		ptr.String("46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1"),
	)
}

var ValidEventNoTags = func() *event.Event {
	return event.NewEvent(
		1,
		"GM nostr welcome to Saturday!",
		nil,
		ptr.Int64(1712350548),
		ptr.String("b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76"),
		ptr.String("234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983"),
		ptr.String("46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1"),
	)
}

var ValidEventMinimal = func() *event.Event {
	return event.NewEvent(
		1,
		"hello world!",
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

var ValidEventJSON = `{"kind":1,"content":"GM nostr welcome to Saturday!","tags":[],"created_at":1712350548,"id":"b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76","pubkey":"234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983","sig":"46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1"}`

var ValidEvent2JSON = `{"kind":1,"content":"#[0]’s desire for more micro apps on nostr is critical. \n\nWe’re having fun with all the social clients being built right now, but the true power of this protocol comes with thousands of smaller ultilities coming together to build an ecosystem of valuable services. The seamlessness of switching between them will be the magic. \n\nI think that’s where this becomes truly unique. Can’t wait to see more.","tags":[["p","3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d"]],"created_at":1673311423,"id":"9007b89f5626b945174a2a8c8d9d0aefc44389fcdd45da2d14ec21bd2f943efe","pubkey":"82341f882b6eabcd2ba7f1ef90aad961cf074af15b9ef44a09f9d2a8fbfbe6a2","sig":"f188ace3426d97dbe1641b35984dc839a5c88a728e7701c848144920616967eb64a30a7d657ca16d556bea718311b15260c886568531399ed14239868aedbcee"}`

const ValidEventSerialize = "[0,\"234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983\",1712350548,1,[],\"GM nostr welcome to Saturday!\"]"
const ValidEventString = `{"kind":1,"content":"GM nostr welcome to Saturday!","tags":null,"created_at":1712350548,"id":"b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76","pubkey":"234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983","sig":"46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1"}`

var InvalidEventCreatedAt = func() *event.Event {
	return event.NewEvent(
		1,
		"GM nostr welcome to Saturday!",
		event.Tags{nil},
		ptr.Int64(-1),
		ptr.String("dd5f42113f06e43ed870a76"),
		ptr.String("f8e13fbe43890f8656ee983"),
		ptr.String("9e1894b5fd09195a65e72df7d9e7c1"),
	)
}

var InvalidEventSignature = func() *event.Event {
	return event.NewEvent(
		1,
		"GM nostr welcome to Saturday!",
		event.Tags{nil},
		ptr.Int64(1),
		ptr.String("dd5f42113f06e43ed870a76dd5f42113f06e43ed870a76dd5f42113f06e43ed8"),
		ptr.String("f8e13fbe43890f8656ee983f8e13fbe43890f8656ee983f8e13fbe43890f8656"),
		ptr.String("9efffffffffffffb4b4b4b4b4be7c19e01234fffffffffffffb4b4b4b4b4be7c9efffffffffffffb4b4b4b4b4be7c19e01234fffffffffffffb4b4b4b4b4be7c"), // hex but invalid for pubkey/ID
	)
}

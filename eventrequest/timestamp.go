package eventrequest

import (
	"time"
)

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0).UTC()
}

func Now() Timestamp {
	return Timestamp(time.Now().Unix())
}

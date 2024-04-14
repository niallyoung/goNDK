package eventrequest_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

func TestTime_Now(t *testing.T) {
	t.Run("Now", func(t *testing.T) {
		now := event.Now()
		assert.Greater(t, now, event.Timestamp(1712660133))
		assert.Greater(t, now.Time(), event.Timestamp(1712660133).Time())
	})
}

func TestTimestamp_Time(t *testing.T) {
	scenarios := []struct {
		name         string
		givenInt64   int64
		expectedTime time.Time
	}{
		{
			"zero",
			0,
			time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"negative",
			-22059601,
			time.Date(1969, time.April, 20, 16, 19, 59, 0, time.UTC),
		},
		{
			"positive",
			2407894000,
			time.Date(2046, time.April, 21, 3, 26, 40, 0, time.UTC),
		},
		{
			"positive small",
			2,
			time.Date(1970, time.January, 1, 0, 0, 2, 0, time.UTC),
		},
		{
			"positive large",
			9481283247,
			time.Date(2270, time.June, 14, 1, 47, 27, 0, time.UTC),
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ts := event.Timestamp(s.givenInt64)
			assert.Equal(t, s.expectedTime, ts.Time())
		})
	}
}

package client

import "github.com/niallyoung/goNDK/event"

// A Filter is a filter for subscription.
type Filter struct {
	IDs     []string   `json:"ids,omitempty"`
	Kinds   []int      `json:"kinds,omitempty"`
	Authors []string   `json:"authors,omitempty"`
	Tags    event.Tags `json:"-,omitempty"`
	Since   int64      `json:"since,omitempty"`
	Until   int64      `json:"until,omitempty"`
	Limit   int        `json:"limit,omitempty"`
	Search  string     `json:"search,omitempty"`
}

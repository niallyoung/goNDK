package client

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Client struct {
	RelayManager map[string]RelayManager
}

func NewClient(urls ...string) Client {
	client := Client{
		RelayManager: make(map[string]RelayManager, len(urls)),
	}

	for _, url := range urls {
		client.RelayManager[url] = *NewRelayManager(url)
	}

	return client
}

func (c Client) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.RelayManager, validation.Each(is.PrintableASCII)),
	)
}

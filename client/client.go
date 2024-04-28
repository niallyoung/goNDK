package client

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Validate() error {
	return validation.ValidateStruct(&c)//validation.Field(&c.Foo, validation.Required),

}

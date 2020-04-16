package goamazon

import (
	"github.com/hunterhug/marmot/miner"
	"time"
)

type Client struct {
	Proxy    bool
	WaitTime time.Duration
}

func NewClient() *Client {
	return &Client{}
}

func New() *Client {
	return NewClient()
}

func (c *Client) SetProxy(proxy bool) *Client {
	c.Proxy = proxy
	return c
}

func (c *Client) SetLogLevel(level string) *Client {
	miner.SetLogLevel(level)
	return c
}

func (c *Client) SetDebug() *Client {
	c.SetLogLevel(miner.DEBUG)
	return c
}

func (c *Client) SetWaitTime(duration time.Duration) *Client {
	if duration <= 0 {
		return c
	}
	c.WaitTime = duration
	return c
}

package goamazon

import "github.com/hunterhug/marmot/miner"

type Client struct {
	Proxy bool
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

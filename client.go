package reddit

import (
	"time"

	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	boards         *gocache.InterfaceCache
	subscribers    *gocache.InterfaceCache
	updateInterval time.Duration
}

func NewClient() *Client {
	return &Client{
		boards:      gocache.NewInterfaceCache(),
		subscribers: gocache.NewInterfaceCache(),
	}
}

func (c *Client) AddListener(key string, ch chan []T3Data) {
	c.subscribers.Set(key, ch)
}

func (c *Client) AddBoard(b string) error {
	return nil
}

//Listen - listens to the selected boards and dispatches changes to the users
func (c *Client) Listen() {
	for {

		time.Sleep(c.updateInterval)
	}
}

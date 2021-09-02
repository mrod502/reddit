package reddit

import (
	"fmt"

	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	subreddits  *gocache.InterfaceCache
	subscribers *gocache.InterfaceCache
	dispatcher  gocache.Dispatcher
}

func NewClient() *Client {
	return &Client{
		subreddits:  gocache.NewInterfaceCache(),
		subscribers: gocache.NewInterfaceCache(),
	}
}

func (c *Client) AddListener(key string, ch chan []T3Data) {
	c.subscribers.Set(key, ch)
}

func (c *Client) AddSub(b string) error {
	v, err := GetSub(b)
	if err != nil {
		return err
	}
	c.subreddits.Set(b, &v)
	return nil
}

func (c *Client) RemoveSub(b string) {
	c.subreddits.Delete(b)
}

func (c *Client) GetSubreddit(b string) (*Subreddit, error) {
	if !c.subreddits.Exists(b) {
		return nil, ErrNotFound
	}
	if v, ok := c.subreddits.Get(b).(*Subreddit); ok {
		return v, nil
	}
	return nil, ErrTypeAssertion
}

func (c *Client) GetPostReplies(boardId, postId string) (a []T3Data, err error) {
	board, ok := c.subreddits.Get(boardId).(*Subreddit)

	if !ok {
		return a, ErrNotFound
	}

	post, err := FindPost(board.Data.Children, func(l Link) bool { return l.Data.Id == postId })

	if err != nil {
		return a, err
	}

	return post.GetReplies()
}

func FindPost(posts []Link, finder func(Link) bool) (Link, error) {
	for _, v := range posts {
		if finder(v) {
			return v, nil
		}
	}
	return Link{}, ErrNotFound
}

//Listen - listens to the selected boards and dispatches changes to the users
func (c *Client) Refresh() {
	for _, k := range c.subreddits.GetKeys() {
		v, ok := c.subreddits.Get(k).(*Subreddit)
		if !ok || (v == nil) {
			continue
		}
		err := v.Update()
		if err != nil {
			fmt.Printf("failed to update %s: %s\n", k, err.Error())
		}
	}

}

func (c *Client) SetDispatchFunc(f gocache.Dispatcher) {
	c.dispatcher = f
}

package reddit

import (
	"time"

	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	subreddits  *gocache.InterfaceCache
	subscribers *gocache.InterfaceCache
	dispatcher  gocache.Dispatcher
}

func NewClient(ttl time.Duration) *Client {
	return &Client{
		subreddits:  gocache.NewInterfaceCache().WithExpiration(ttl),
		subscribers: gocache.NewInterfaceCache(),
	}
}

func (c *Client) AddListener(key string, ch chan []T3Data) {
	c.subscribers.Set(key, ch)
}

func (c *Client) RemoveSub(b string) {
	c.subreddits.Delete(b)
}

func (c *Client) GetSubreddit(b string) (*Subreddit, error) {
	if !c.subreddits.Exists(b) {
		s, err := getSub(b)
		c.subreddits.Set(b, s)
		return s, err
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

	post, err := FindPost(board.Data.Children, func(l *Link) bool { return l.Data.Id == postId })

	if err != nil {
		return a, err
	}

	return post.GetReplies()
}

func FindPost(posts []*Link, finder func(*Link) bool) (*Link, error) {
	for _, v := range posts {
		if finder(v) {
			return v, nil
		}
	}
	return &Link{}, ErrNotFound
}

func (c *Client) SetDispatchFunc(f gocache.Dispatcher) {
	c.dispatcher = f
}

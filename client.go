package reddit

import (
	"time"

	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	subreddits  *gocache.InterfaceCache
	subscribers *gocache.InterfaceCache
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
	if v, err := c.subreddits.Get(b); err == nil {
		return v.(*Subreddit), nil
	} else {
		return nil, ErrTypeAssertion
	}
}

func (c *Client) GetPostReplies(boardId, postId string) (a []T3Data, err error) {
	board, err := c.subreddits.Get(boardId)
	if err != nil {
		return nil, err
	}

	sub := board.(*Subreddit)
	post, err := FindPost(sub.Data.Children, func(l *Link) bool { return l.Data.Id == postId })

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

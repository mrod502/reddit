package reddit

import (
	"encoding/json"
	"fmt"
	"time"
)

type Subreddit struct {
	Data Data   `json:"data"`
	Kind string `json:"kind"`
	sr   string
}

func (b Subreddit) GetComments() (c []T3Data) {
	c = make([]T3Data, len(b.Data.Children))
	for i, v := range b.Data.Children {
		c[i] = v.Data
	}
	return
}

func (b Subreddit) GetHrefs() (o []string) {
	for _, v := range b.Data.Children {
		o = append(o, v.GetHrefs()...)
	}
	return
}

func (b Subreddit) GetPostReplies(post T3Data) (d []*Link, err error) {
	bytes, _, err := BrowserRequest(fmt.Sprintf("%sr/%s/comments/%s.json", redditBase, post.Subreddit, post.Id))
	if err != nil {
		return
	}
	var l ListingArray

	err = json.Unmarshal(bytes, &l)

	d = l.AllChildren()
	return
}

func (b Subreddit) Subreddit() string {
	if b.sr != "" {
		return b.sr
	}

	if len(b.Data.Children) == 0 {
		return ""
	}
	b.sr = b.Data.Children[0].Data.Subreddit
	return b.sr
}

func (b Subreddit) GetAllDiscussion() (t []T3Data, waitTime time.Duration, err error) {
	childIDs := b.Data.ChildIDs()

	for _, v := range childIDs {
		listings, er := GetCommentListing(b.Subreddit(), v, "?sort=top")

		if er != nil {
			err = er
		}
		for _, child := range listings.AllChildren() {
			t = append(t, child.Data.GetStockSymbols())
		}

		time.Sleep(time.Second + time.Second/4)

	}
	return
}

func (b *Subreddit) Update() error {
	sub, err := getSub(b.Subreddit())

	if err != nil {
		return err
	}
	b.Data = sub.Data
	return nil

}

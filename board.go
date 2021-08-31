package reddit

import (
	"encoding/json"
	"fmt"
)

type Board struct {
	Data Data   `json:"data"`
	Kind string `json:"kind"`
}

func (b Board) GetComments() (c []T3Data) {
	c = make([]T3Data, len(b.Data.Children))
	for i, v := range b.Data.Children {
		c[i] = v.Data
	}
	return
}

func (b Board) GetHrefs() (o []string) {
	for _, v := range b.Data.Children {
		o = append(o, v.GetHrefs()...)
	}
	return
}

func (b Board) GetPostReplies(post T3Data) (d []T1Data, err error) {
	bytes, _, err := BrowserRequest(fmt.Sprintf("%sr/%s/comments/%s.json", redditBase, post.Subreddit, post.ID))
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &d)
	return
}

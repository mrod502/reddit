package reddit

import "time"

type Link struct {
	Kind        string `json:"kind,omitempty"`
	Data        T3Data `json:"data"`
	Replies     []T3Data
	lastUpdated time.Time
}

func (l Link) GetHrefs() (o []string) {
	matches := mdLinkRegex.FindAllStringSubmatch(l.Data.Selftext, -1)
	for _, m := range matches {
		if len(m) == 3 {
			o = append(o, m[2])
		}
	}
	return o
}

func (l Link) IsDailyDiscussion() bool {
	if len(l.Data.LinkFlairRichText) == 0 {
		return false
	}
	return l.Data.LinkFlairRichText[0].T == "Daily Discussion"
}
func (r Data) ChildIDs() (ids []string) {
	ids = make([]string, 0, len(r.Children))
	for _, val := range r.Children {
		ids = append(ids, val.Data.Id)
	}
	return ids
}

func (l *Link) GetReplies(expiration ...time.Duration) ([]T3Data, error) {
	var err error
	if len(expiration) == 0 {
		l.Replies, err = l.Data.GetReplies()
		return l.Replies, err
	}

	if tnow := time.Now(); expiration[0] < tnow.Sub(l.lastUpdated) {
		l.lastUpdated = tnow

		l.Replies, err = l.Data.GetReplies()
		return l.Replies, err
	}

	return l.Replies, nil
}

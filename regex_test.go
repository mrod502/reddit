package reddit

import (
	"fmt"
	"testing"
)

func TestLinkRegex(t *testing.T) {
	var strContainingLink = `Your daily trading discussion thread. Please keep the shitposting to a minimum. \n\n^Navigate ^WSB |^We ^recommend ^best ^daily ^DD\n:--|:--                                 \n**DD** | [All](https://reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3ADD) / [**Best Daily**](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ADD&amp;restrict_sr=on&amp;t=day) / [Best Weekly](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ADD&amp;restrict_sr=on&amp;t=week)\n**Discussion** | [All](https://reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3ADiscussion) / [**Best Daily**](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ADiscussion&amp;restrict_sr=on&amp;t=day) / [Best Weekly](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ADiscussion&amp;restrict_sr=on&amp;t=week)\n**YOLO** | [All](https://reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AYOLO) / [**Best Daily**](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3AYOLO&amp;restrict_sr=on&amp;t=day) / [Best Weekly](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3AYOLO&amp;restrict_sr=on&amp;t=week)\n**Gain** | [All](https://reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AGain) / [**Best Daily**](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3AGain&amp;restrict_sr=on&amp;t=day) / [Best Weekly](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3AGain&amp;restrict_sr=on&amp;t=week)\n**Loss** | [All](https://reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3ALoss) / [**Best Daily**](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ALoss&amp;restrict_sr=on&amp;t=day) / [Best Weekly](https://www.reddit.com/r/wallstreetbets/search?sort=top&amp;q=flair%3ALoss&amp;restrict_sr=on&amp;t=week)\n\n\n[Weekly Earnings Discussion Thread](https://www.reddit.com/r/wallstreetbets/search?sort=new&amp;restrict_sr=on&amp;q=flair%3A%22Earnings%20Thread%22)\n\n**Read the [rules](https://www.reddit.com/r/wallstreetbets/wiki/contentguide) and make sure other people follow them.**\n\nTry [No Meme Mode](https://www.reddit.com/r/wallstreetbets/search/?q=-flair%3AMeme%20-flair%3ASatire%20-flair%3AShitpost&amp;restrict_sr=1&amp;t=day&amp;sort=hot), also accessible through the top bar.\n\nFollow [@Official_WSB](https://twitter.com/Official_WSB) on Twitter, all other accounts are impersonators.\n\nCheck out our [Discord](https://discord.gg/wallstreetbets)`

	matches := mdLinkRegex.FindAllStringSubmatch(strContainingLink, -1)
	for _, v := range matches {
		fmt.Println(v[2])
	}
}

func TestSymbolRegex(t *testing.T) {
	var strContainingSymols = `$GME to the #MOON`
	var want = [][]string{{"$GME", "$", "GME"}, {"#MOON", "#", "MOON"}}
	have := symbolRegex.FindAllStringSubmatch(strContainingSymols, -1)
	if fmt.Sprint(want) != fmt.Sprint(have) {
		t.Fatalf("have:%s\nwant:%s", have, want)
	}
}

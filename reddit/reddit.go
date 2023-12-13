package reddit

import (
    "strings"
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "context"
    "errors"
)

var ctx = context.Background()

func GetPosts() ([]string, error) {
    var allowed []string = []string{".jpg", ".jpeg", ".png", ".gif"}
	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, "lotrmemes", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "day",
	})
	if err != nil {
		return nil, errors.New("Could not get posts")
	}
    var filteredPosts []string
	for _, post := range posts {
        for i := 0; i < len(allowed); i++ {
            if strings.HasSuffix(post.URL, allowed[i]) {
                filteredPosts = append(filteredPosts, post.URL)
            }
        }
	}
	return filteredPosts, nil
}

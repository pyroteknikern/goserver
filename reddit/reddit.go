package reddit

import (
    "strings"
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "context"
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

var ctx = context.Background()

func GetPosts() ([]string, error) {
    var allowed []string = []string{".jpg", ".jpeg", ".png", ".gif"}
    
    godotenv.Load()

    secret := os.Getenv("SECRET")
    id := os.Getenv("ID")
    username := os.Getenv("USERNAME")
    password := os.Getenv("PASSWORD")

    fmt.Println(id)
    credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}
    posts,_ , err := client.Subreddit.TopPosts(ctx, "lotrmemes", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "day",
	})
	if err != nil {
		return nil, err
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

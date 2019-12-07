package website

import (
	"composition/post"
	"fmt"
)

type Website struct {
	Posts []post.Post
}

func (website *Website) ListAllPosts() {
	for _, v := range website.Posts {
		v.Details()
		fmt.Println()
	}
}

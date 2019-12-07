package post

import (
	"composition/author"
	"fmt"
)

type Post struct {
	Title, Content string
	author.Author
}

func (post *Post) Details() {
	fmt.Println("Title :", post.Title)
	fmt.Println("Content :", post.Content)
	fmt.Println("Author :", post.FullName())
	fmt.Println("Bio :", post.Bio)
}

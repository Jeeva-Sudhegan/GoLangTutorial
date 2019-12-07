package main

import (
	"composition/author"
	"composition/post"
	"composition/website"
)

/*
go not support inheritance but support composition
composition is achieved by embedding structs inside structs
*/

func main() {
	author1 := author.Author{
		"Jeeva Sudhegan",
		"G",
		"Budding programmer",
	}
	post1 := post.Post{
		"The Legend1",
		"Story about the lost Sword1",
		author1,
	}
	post2 := post.Post{
		"The Legend2",
		"Story about the lost Sword2",
		author1,
	}
	post3 := post.Post{
		"The Legend3",
		"Story about the lost Sword3",
		author1,
	}
	website1 := website.Website{
		Posts: []post.Post{post1, post2, post3},
	}
	website1.ListAllPosts()
}

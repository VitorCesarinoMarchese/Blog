package Blogpost

import (
	"blog.net/db"
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title       string `json:"Title"`
	Img         string `json:"Img"`
	Description string `json:"Description"`
	Author      string `json:"Author"`
	Views       int `json:"Views"`
}

func CreatePost(post *BlogPost) {
	daba := db.ConnectDb()
	daba.AutoMigrate(&BlogPost{})
	result := daba.Create(post)
	if result.Error != nil {
		panic(result.Error)
	}
}
func FindPostById(id uint) BlogPost{
	var Post BlogPost
	daba := db.ConnectDb()
	daba.First(&Post, id)
	return Post
}
func ReadPosts() []BlogPost {
	daba := db.ConnectDb()
	var BlogPosts []BlogPost
	daba.Find(&BlogPosts)
	return BlogPosts
}

package Blogpost

import (
	"blog.net/db"
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title, Img, Description, Author string
	Views                           int
}
func CreatePost(post BlogPost){
	db := db.ConnectDb()
	db.AutoMigrate(&BlogPost{})
	result := db.Create(post)
	if result.Error != nil{
		panic(result.Error)
	}
}
func ReadPosts() []BlogPost{
	db := db.ConnectDb()
	var BlogPosts []BlogPost
	db.Find(&BlogPosts)
	return BlogPosts
}
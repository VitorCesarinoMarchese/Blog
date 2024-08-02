package main

import (
	m "blog.net/Models"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.POST("/createPost", func(c *gin.Context) {
		var newPost m.BlogPost
		if err := c.BindJSON(&newPost); err != nil{
			panic(err)
		}
		m.CreatePost(newPost)
	})
	r.GET("/Posts", func(c *gin.Context){
		c.JSON(200, m.ReadPosts())
	})


	r.Run(":8000")
}

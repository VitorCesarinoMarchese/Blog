package main

import (
	"strconv"

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
	r.GET("/post/:id", func(c *gin.Context) {
		id := c.Param("id")
		nid, err := strconv.Atoi(id)
		if err != nil{
			panic(err)
		}
		c.JSON(200, m.FindPostById(uint(nid)))
	})
	r.POST("/createPost", func(c *gin.Context) {
		newPost := &m.BlogPost{}
		if err := c.BindJSON(&newPost); err != nil{
			panic("err")
		}
		m.CreatePost(newPost)
		result := m.FindPostById(newPost.ID)
		c.JSON(200, result)
	})
	r.GET("/Posts", func(c *gin.Context){
		c.JSON(200, m.ReadPosts())
	})
	r.Run(":8000")
}

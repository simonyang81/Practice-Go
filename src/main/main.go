package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	name := "Simon"
	fmt.Println("Hello, " + name)

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Print(s)
	var u user
	u.name = "Simon Y"
	u.email = "simonist.yang@gmail.com"
	u.changeEmail("simonist.yang@outlook.com")

	fmt.Print(u.email)

	//initGin()
}

func initGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin",
		})
	})

	r.Run("localhost:8085")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	router.Static("/assets", "./assets/")
	router.StaticFile("/", "./index.html")
	router.Run(":8080")
}

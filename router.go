package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l-f-h/judge_server/method"
	"log"
)

func main() {
	router := gin.Default()
	router.POST("/question/judge", method.Judge)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Gin Run failed")
	}
}

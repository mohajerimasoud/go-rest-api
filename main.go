package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/controller"
	"log"
)

func appRouter() {
	router := gin.Default()
	router.GET("/articles", controller.GetAllArticles)
	router.POST("/articles", controller.CreateArticle)

	log.Fatal(router.Run(":8080"))
}

func main() {
	appRouter()
}

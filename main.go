package main

import (
	"links_available/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/check_available_links", handlers.CheckAvailableLinks)
		api.POST("/get_links_list", handlers.GetSavedLinks)
	}

	r.Run(":8082")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Serve static files from the "static" directory
	//r.Use(static.Serve("/", static.LocalFile("./static", true)))
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Route to serve the HTML file
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Rendering home page")
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	// Run the server
	r.Run(":8080")

}

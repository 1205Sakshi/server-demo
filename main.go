package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting main method")
	// Initialize a new Gin router
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		fmt.Println("Test router")
		c.JSON(http.StatusOK, gin.H{
			"output": "Working!",
		})
	})

	// Run the Gin server on port 8080
	if err := router.Run(":8089"); err != nil {
		// If there's an error starting the server, print the error
		fmt.Println("Error starting server:", err)
	}
}

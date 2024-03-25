// This sets up a server with a single endpoint that sends a server-sent events
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", increments)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func increments(c *gin.Context) {
	total := 10000
	index := 0

	for index <= total {

		c.SSEvent("index", map[string]interface{}{
			"Index": index,
			"Total": total,
		})
		c.Writer.Flush()
		index += 1
	}

	c.Writer.Flush()
}

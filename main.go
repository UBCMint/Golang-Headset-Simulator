// This sets up a server with a single endpoint that sends a server-sent events
package main

import (
	"math/rand"

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

	var ContinueForwards bool = true

	for ContinueForwards {
		matrix := generateRandomMatrix()

		c.SSEvent("index", map[string]interface{}{
			"Index":  index,
			"Total":  total,
			"Matrix": matrix,
		})
		c.Writer.Flush()
		index += 1
	}

	c.Writer.Flush()
}

func generateRandomMatrix() [][]int {
	matrix := make([][]int, 10)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 10)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	return matrix
}

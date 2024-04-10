// This sets up a server with a single endpoint that sends a server-sent events
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	println("Server is running on port http://localhost:8080")
	http.HandleFunc("/headsetData", eventsHandler)
	http.ListenAndServe(":8080", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var ContinueForwards bool = true

	fmt.Fprintf(w, "Data is being sent\n\n")

	for ContinueForwards {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", generateRandomMatrix()))
		w.(http.Flusher).Flush()
	}

	closeNotify := r.Context().Done()
	<-closeNotify
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

// Alternative way to do this with gin
// import (
// 	"math/rand"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/", increments)

// 	// Start the server
// 	if err := router.Run(":8080"); err != nil {
// 		panic(err)
// 	}
// }

// func increments(c *gin.Context) {
// 	total := 10000
// 	index := 0

// 	var ContinueForwards bool = true

// 	for ContinueForwards {
// 		matrix := generateRandomMatrix()

// 		c.SSEvent("index", map[string]interface{}{
// 			"Index":  index,
// 			"Total":  total,
// 			"Matrix": matrix,
// 		})
// 		c.Writer.Flush()
// 		index += 1
// 	}

// 	c.Writer.Flush()
// }

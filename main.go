package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v1",
			"data":    "masuk",
		})
	})

	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		fmt.Println("Listen and Serve at port 8000")
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()

}

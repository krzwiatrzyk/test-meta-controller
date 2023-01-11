package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/labeler", labelerHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	r.Run(fmt.Sprintf(":%s", port))

// 	// Shutdown on SIGTERM.
// 	sigchan := make(chan os.Signal, 2)
// 	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
// 	sig := <-sigchan
// 	log.Printf("Received %v signal. Shutting down...", sig)
}

func labelerHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		println("JSON could not be retrieved")
		c.String(http.StatusBadRequest, "JSON could not be retrieved")
		return
	}
	fmt.Printf("%+v",string(body))
	c.String(http.StatusOK,"Success")
}
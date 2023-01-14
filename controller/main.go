package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var labels = map[string]string{}

func main() {
	r := gin.Default()
	r.POST("/labeler", labelerHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9091"
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
	bodyString := string(body)
	// os.WriteFile("messages/" + time.Now().String() + ".json", body, 0644)
	kind := gjson.Get(bodyString, "object.kind").String()
	
	switch kind {
	case "Label":
		label(c, bodyString)
	case "Deployment":
		deployment(c, bodyString)
	}
}

// func test() {
// 	message, _ := os.ReadFile("../message.json")
// 	bodyString := string(message)
// 	fmt.Println(bodyString)
// 	value := gjson.Get(bodyString, "finalizing")
// 	fmt.Printf("%T - %v", value, value)
// }

func label(c *gin.Context, body string) {
	fmt.Println("label handler")
	label := gjson.Get(body, "object.spec.label").String()
	labels[label] = label
	fmt.Println("Registered label - " + label)
	c.String(http.StatusOK, "Success")
}

func deployment(c *gin.Context, body string) {
	fmt.Println("deployment handler")
	fmt.Println(labels)
	fmt.Println(json.Marshal(response{labels}))
	c.JSON(http.StatusOK, response{labels})
}

type response struct {
	Labels map[string]string `json:"labels"`
}
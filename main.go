package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"clustern-node/shared"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
)

func main() {
	r := gin.Default()
	r.POST("/create_app", func(c *gin.Context) {
		var input shared.CreateDeployedAppInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": "unable to process json input",
			})
			return
		}

		if err := configureApp(input.RepoName, input.Port); err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"message": "Your app has been configured! Beginning deploy...",
		})
	})

	r.GET("/app_status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "Status update",
			"status":  "deploying...",
		})
	})

	fmt.Println("Server has been configured!")
	r.Run(":1337")
}

func configureApp(repoName string, port int) error {
	var logBuffer bytes.Buffer
	repoName = strings.TrimPrefix(repoName, "/")

	_, err := git.PlainClone("../app", false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://github.com/%s", repoName),
		Progress: &logBuffer,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("../app/out.log", logBuffer.Bytes(), 0755); err != nil {
		return err
	}

	return nil
}

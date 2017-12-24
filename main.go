package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})

	})

	router.GET("/resources/:img", func(c *gin.Context) {
		img := fmt.Sprintf("resources/%s", c.Param("img"))

		c.File(img)
	})

	router.GET("/singers/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "singer: %s", name)
	})

	router.GET("/songs/:name", func(c *gin.Context) {
		name := c.Param("name")
		img := fmt.Sprintf("/resources/%s.png", name)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"song":        "成都",
			"keywords":    "成都,赵雷,吉他谱",
			"description": "成都,赵雷,吉他谱",
			"img":         img,
		})
	})

	router.GET("/albums/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "album: %s", name)
	})

	router.Run()

}

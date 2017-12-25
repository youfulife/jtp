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
		//img := fmt.Sprintf("/resources/%s.png", name)
		cdn := "http://ofphfz209.bkt.clouddn.com/"
		img := fmt.Sprintf("%s%s1.png", cdn, name)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"song":   "成都",
			"singer": "赵雷",
			"tags":   "深蓝雨",
			"img":    img,
		})
	})

	router.GET("/albums/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "album: %s", name)
	})

	router.Run()

}

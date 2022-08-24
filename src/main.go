package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		url := c.Query("url")
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": "-1", "message": fmt.Sprintf("[%s] unavailable", url), "data": nil})
		} else {
			defer response.Body.Close()
			c.DataFromReader(
				http.StatusOK,
				response.ContentLength,
				response.Header.Get("Content-Type"),
				response.Body,
				map[string]string{"Content-Disposition": fmt.Sprintf("attachment; filename=%s", path.Base(url))},
			)
		}
	})

	router.Run(":8080")
}

package main

import (
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
			c.Status(http.StatusServiceUnavailable)
			return
		}

		defer response.Body.Close()

		extraHeaders := map[string]string{"Content-Disposition": "attachment; filename=" + path.Base(url)}

		c.DataFromReader(http.StatusOK, response.ContentLength, response.Header.Get("Content-Type"), response.Body, extraHeaders)
	})

	router.Run(":8080")
}

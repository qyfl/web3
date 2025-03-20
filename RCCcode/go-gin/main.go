package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/thinkerou/favicon"

func main() {
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./img.jpg"))
	// Setup Security Headers

	expectedHost := "localhost:8080"
	ginServer.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		cookie, err := context.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  cookie,
		}

		fmt.Println("cookie", cookie)
		context.AsciiJSON(http.StatusOK, data)

		//context.JSON(200, gin.H{"msg": "hello,world"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post request"})
	})

	ginServer.Run()

}

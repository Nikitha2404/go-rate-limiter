package main

import (
	"go-rate-limiter/config"
	"go-rate-limiter/ratelimiter"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfigurations("./config")

	rateLimiter := ratelimiter.InitRateLimiter()
	router := gin.Default()
	router.GET("/hello", middlewareRateLimiter(rateLimiter), messageHandler)

	router.Run(config.AppConfig.Server.Host)
}

func middlewareRateLimiter(r *ratelimiter.RateLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Print("executing rate limiter middleware")
		if !r.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too may requests.Please try later.",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func messageHandler(c *gin.Context) {
	c.Writer.Write([]byte("server accepting requests!"))
}

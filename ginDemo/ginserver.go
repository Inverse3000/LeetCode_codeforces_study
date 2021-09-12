package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	var rid int
	r.Use(func(c *gin.Context) {
		//path, log latency, response code path
		s := time.Now()
		c.Next()
		fmt.Println("???")
		rid = rand.Int()
		logger.Info("incoming request", zap.String("path", c.Request.URL.Path), zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)), zap.Int("RequestId", rid))
	},
		func(c *gin.Context) {
			fmt.Println("i know")
			c.Set("RequestId", rid)
			c.Next()
		})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get("RequestId"); exists {
			h["RequestId"] = rid
		}
		c.JSON(200, h)
	})
	r.Run()
}

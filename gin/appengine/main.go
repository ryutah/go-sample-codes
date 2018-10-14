package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const appengineContext = "appengine_context"

func init() {
	r := gin.New()

	r.Use(func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		c.Set(appengineContext, ctx)
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		val, _ := c.Get(appengineContext)
		ctx := val.(context.Context)

		log.Infof(ctx, "receive request...")
		c.Writer.WriteString("hello world!!")
	})

	http.Handle("/", r)
}

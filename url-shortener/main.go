package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "url-shortener/handler"
)

func main() {
    gin.SetMode(gin.ReleaseMode)

    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())
    r.SetTrustedProxies(nil)

    r.Static("/static", "./static")

    r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    r.GET("/favicon.ico", func(c *gin.Context) {
        c.String(http.StatusNoContent, "")
    })

    r.POST("/shorten", handler.ShortenURL)
    r.GET("/:shortURL", handler.RedirectURL)

    log.Println("Server running on http://0.0.0.0:3000")
    if err := r.Run(":3000"); err != nil {
        log.Fatal("Unable to start:", err)
    }
}

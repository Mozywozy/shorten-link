package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "url-shortener/handler"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // default port for Vercel
    }

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

    log.Println("Server running on port " + port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Unable to start:", err)
    }
}

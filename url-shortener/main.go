package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "url-shortener/handler"
)

func main() {
    gin.SetMode(gin.ReleaseMode) // Set mode produksi

    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery()) // Menambahkan middleware logger dan recovery
    r.SetTrustedProxies(nil) // Percayai semua proxy (untuk pengembangan)

    // Menyajikan file statis
    r.Static("/static", "./static")

    // Tambahkan root endpoint
    r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    // Tambahkan rute untuk favicon
    r.GET("/favicon.ico", func(c *gin.Context) {
        c.String(http.StatusNoContent, "")
    })

    r.POST("/shorten", handler.ShortenURL)
    r.GET("/:shortURL", handler.RedirectURL)

    log.Println("Server running on http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Unable to start:", err)
    }
}

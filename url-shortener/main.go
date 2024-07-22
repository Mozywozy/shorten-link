package main

import (
    "log"
    "net/http"
    "os"

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

    // Gunakan PORT dari variabel lingkungan
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port untuk pengembangan lokal
    }

    log.Printf("Server running on http://localhost:%s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Unable to start:", err)
    }
}

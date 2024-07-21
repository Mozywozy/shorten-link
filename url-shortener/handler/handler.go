package handler

import (
	"log"
	"net/http"
	"url-shortener/model"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
    var request model.URLRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    shortURL := store.SaveURL(request.LongURL)
    c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func RedirectURL(c *gin.Context) {
    shortURL := c.Param("shortURL")
    log.Printf("Received short URL: %s", shortURL)

    longURL, err := store.GetURL("http://localhost:8080/" + shortURL) // Pastikan formatnya sesuai
    if err != nil {
        log.Printf("URL not found for %s", shortURL)
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    log.Printf("Redirecting %s to %s", shortURL, longURL)
    c.Redirect(http.StatusMovedPermanently, longURL)
}

package store

import (
    "crypto/rand"
    "errors"
    "fmt"
    "sync"
)

var (
    urlStore = make(map[string]string)
    mu       sync.Mutex
)

const (
    shortURLLength = 6
    charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateRandomString(length int) string {
    b := make([]byte, length)
    _, err := rand.Read(b)
    if err != nil {
        return ""
    }

    for i := range b {
        b[i] = charset[int(b[i])%len(charset)]
    }
    return string(b)
}

func SaveURL(longURL string) string {
    mu.Lock()
    defer mu.Unlock()

    var shortURL string
    for {
        shortURL = fmt.Sprintf("http://localhost:8080/%s", generateRandomString(shortURLLength))
        if _, exists := urlStore[shortURL]; !exists {
            break
        }
    }

    urlStore[shortURL] = longURL
    return shortURL
}

func GetURL(shortURL string) (string, error) {
    mu.Lock()
    defer mu.Unlock()

    longURL, exists := urlStore[shortURL]
    if !exists {
        return "", errors.New("URL not found")
    }
    return longURL, nil
}

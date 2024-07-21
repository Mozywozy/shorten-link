package model

type URLRequest struct {
    LongURL string `json:"long_url" binding:"required"`
}

package req

import (
	"net/http"
	"time"
)

// HTTPClient interface for test purpose
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// GetClient to get a client with timeout info
func GetClient(timeout int) *http.Client {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &client
}

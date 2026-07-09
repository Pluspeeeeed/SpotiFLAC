package backend

import (
	"net/http"
	"strings"
	"time"
)

type LyricsClient struct {
	httpClient *http.Client
}

func NewLyricsClient() *LyricsClient {
	return &LyricsClient{
		httpClient: NewHTTPClient(15 * time.Second),
	}
}

// ... rest of the file remains the same

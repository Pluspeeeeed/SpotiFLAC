package backend

import (
	"net"
	"net/http"
	"time"
)

// NewHTTPClient creates an HTTP client that respects system proxy settings.
// On Windows, it reads from environment variables and WinHTTP settings.
// It uses the specified timeout for the overall request.
func NewHTTPClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

// NewHTTPClientWithoutTimeout creates an HTTP client that respects system proxy settings
// without a request timeout. Use sparingly for long-running operations.
func NewHTTPClientWithoutTimeout() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: transport,
	}
}

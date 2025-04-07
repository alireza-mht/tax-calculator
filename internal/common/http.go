package common

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpRequestWithResponse makes a request to the provided url and returns the response object
func HttpRequestWithResponse(httpUrl, method, username, password string) (resp *http.Response, err error) {
	client := &http.Client{}

	httpUrlLowerCase := strings.ToLower(httpUrl)
	if !(strings.HasPrefix(httpUrlLowerCase, "http://") || strings.HasPrefix(httpUrlLowerCase, "https://")) {
		httpUrl = "http://" + httpUrl // Default to http for this project
	}

	// Create empty body or specified body
	var body io.Reader
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		// Use empty string body for methods that expect a body to avoid hanging issues
		body = strings.NewReader("{}")
	} else {
		body = http.NoBody
	}

	// Create request with the specified body
	req, err := http.NewRequest(method, httpUrl, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request: %w", err)
	}

	// Add authentication if provided
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	// Set common headers
	req.Header.Add("Content-Type", "application/json")

	// Execute the request
	return client.Do(req)
}

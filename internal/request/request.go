package request

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// create a http client with timeout of 10 seconds to prevent hanging requests

var client = &http.Client{
	Timeout: 10 * time.Second,
}

// function for resuable request handler also add headers for contentType and accept as application/json

func RequestHandler(method string, url string, body io.Reader) (*http.Response, error) {

	// creating a req  with the method url and body

	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// fmt.Printf(req.Method)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// execute the request

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

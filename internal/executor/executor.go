package executor

import (
	"io"
	"net/http"
	"pmc/internal/request"
	"pmc/internal/response"
	"time"
)

func Executor(method string, url string, bodyReader io.Reader, headerData map[string]string) (string, string, http.Header, string, int64, int, time.Duration) {
	resp, err, elapsed := request.RequestHandler(method, url, bodyReader, headerData)
	body, status, headerDetails, responsemethod, contentLength, statusCode, timeTaken := response.ResponseHandler(resp, err, elapsed)
	return body, status, headerDetails, responsemethod, contentLength, statusCode, timeTaken
}

package executor

import (
	"io"
	"net/http"
	"pmc/internal/request"
	"pmc/internal/response"
	"time"
)

type ResultType struct {
	Body           string
	Status         string
	HeaderDetails  http.Header
	Responsemethod string
	ContentLength  int64
	StatusCode     int
	TimeTaken      time.Duration
}

func Executor(method string, url string, bodyReader io.Reader, headerData map[string]string) ResultType {
	resp, err, elapsed := request.RequestHandler(method, url, bodyReader, headerData)
	body, status, headerDetails, responsemethod, contentLength, statusCode, timeTaken := response.ResponseHandler(resp, err, elapsed)
	result := ResultType{
		Body:           body,
		Status:         status,
		HeaderDetails:  headerDetails,
		Responsemethod: responsemethod,
		ContentLength:  contentLength,
		StatusCode:     statusCode,
		TimeTaken:      timeTaken,
	}
	return result
}

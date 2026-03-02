package executor

import (
	"io"
	"pmc/internal/request"
	"pmc/internal/response"
)

func Executor(method string, url string, bodyReader io.Reader, headerData map[string]string) {
	resp, err, elapsed := request.RequestHandler(method, url, bodyReader, headerData)
	response.ResponseHandler(resp, err, elapsed)
}

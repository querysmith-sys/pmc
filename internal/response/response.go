package response

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ResponseHandler(resp *http.Response, err error, timeTaken time.Duration) {

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	statusCode := resp.StatusCode
	status := resp.Status
	headerDetails := resp.Header
	method := resp.Request.Method
	ContentLength := resp.ContentLength
	if err != nil {
		fmt.Printf("Error Reading response")
		return
	}
	fmt.Printf("Response: %s\n", string(data))
	fmt.Printf("Status: %s\n", status)
	fmt.Printf("HeaderDetails: %s\n", headerDetails)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Content-Length: %d\n", ContentLength)
	fmt.Printf("statusCode: %d\n", statusCode)
	fmt.Printf("TimeTaken: %s\n", timeTaken)
}

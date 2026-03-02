package response

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ResponseHandler(resp *http.Response, err error, timeTaken time.Duration) (string, string, http.Header, string, int64, int, time.Duration) {

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", "", nil, "", 0, 0, 0
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	statusCode := resp.StatusCode
	status := resp.Status
	headerDetails := resp.Header
	responsemethod := resp.Request.Method
	ContentLength := resp.ContentLength
	if err != nil {
		fmt.Printf("Error Reading response")
		return "", "", nil, "", 0, 0, 0
	}
	body := string(data)
	//  i will be returning this all instead
	return body, status, headerDetails, responsemethod, ContentLength, statusCode, timeTaken
	// fmt.Printf("Response: %s\n", string(data))
	// fmt.Printf("Status: %s\n", status)
	// fmt.Printf("HeaderDetails: %s\n", headerDetails)
	// fmt.Printf("Method: %s\n", method)
	// fmt.Printf("Content-Length: %d\n", ContentLength)
	// fmt.Printf("statusCode: %d\n", statusCode)
	// fmt.Printf("TimeTaken: %s\n", timeTaken)
}

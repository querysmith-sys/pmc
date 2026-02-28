package bodybuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var headerData map[string]string
var bodyReader io.Reader

func Bodybuilder(filePath string, header string, body string) (io.Reader, map[string]string) {

	// looping and storing to headerdata
	json.Unmarshal([]byte(header), &headerData)
	if headerData == nil {
		headerData = make(map[string]string)
	}

	// checking if user set the content type or not
	hasContentType := false
	for k := range headerData {
		if strings.ToLower(k) == "content-type" {
			hasContentType = true
			break
		}
	}

	// choice, if file is given use formdata for request if not use json
	if filePath != "" {
		file, err := os.Open(filePath)
		fmt.Printf("Successfully opened %s for upload!\n", filePath)
		if err != nil {
			fmt.Printf("Error Opening File: %v\n", err)
			return nil, nil
		}
		defer file.Close()

		// we will transform raw stream of data to form data format
		buf := &bytes.Buffer{}             // a container for formdata
		writer := multipart.NewWriter(buf) // this  start writting the formdata structure such as boundary.
		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			fmt.Printf("Error creating form file: %v\n", err)
			return nil, nil
		}
		_, err = io.Copy(part, file) // copied the raw bytes from file to the formdata
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
			return nil, nil
		}

		if body != "" {
			var bodyMap map[string]interface{}
			err := json.Unmarshal([]byte(body), &bodyMap)
			if err != nil {
				log.Fatal(err)
			}
			for k, v := range bodyMap {
				writer.WriteField(k, fmt.Sprintf("%v", v))
			}
		}
		writer.Close()
		bodyReader = buf // body contains the form data
		if !hasContentType {
			headerData["Content-Type"] = writer.FormDataContentType()
		}
	} else {
		if body != "" {
			bodyReader = strings.NewReader(body)
			if !hasContentType {
				headerData["Content-Type"] = "application/json"
			}
		} else {
			fmt.Printf("Body is Empty.")
		}

	}

	return bodyReader, headerData
}

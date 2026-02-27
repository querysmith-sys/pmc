package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"pmc/internal/request"
	"strings"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "POST [url]",
	Short: "Send a HTTP POST request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending POST request to: %s\n", url)
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			fmt.Printf("Error File Path Empty")
		}
		file, error := os.Open(filePath)
		if error != nil {
			fmt.Printf("Error Opening File: %v\n", error)
			return
		}
		defer file.Close()

		fmt.Printf("Successfully opened %s for upload!\n", filePath)
		var headerData map[string]string
		err := json.Unmarshal([]byte(header), &headerData)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := request.RequestHandler("POST", url, strings.NewReader(body), headerData)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Error Reading response")
			return
		}
		fmt.Printf("Response: %s\n", string(data))
	},
}

func init() {
	// create a flag for body
	postCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the post request")
	// defineing a  flag for custom header in the format of key and value and mulyiple header can be added by seperating them with comma and key and value should be seperated by colon
	postCmd.Flags().StringP("header", "H", "", "Provide custom headers in the format key:value,key:value")
	// defining a flag for  multipart/formData file upload
	postCmd.Flags().StringP("file", "-F", "", "Flag for uploading File")
	Rootcmd.AddCommand(postCmd)
}

//  -H "content-Type":"json/content-Type","key2":"value2","key3":"value2", "key4":"value2"

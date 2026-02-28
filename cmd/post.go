package cmd

import (
	"fmt"
	"io"
	bodybuilder "pmc/internal/body-builder"
	"pmc/internal/request"

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

		bodyReader, headerData := bodybuilder.Bodybuilder(filePath, header, body)

		resp, err := request.RequestHandler("POST", url, bodyReader, headerData)

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
	postCmd.Flags().StringP("file", "F", "", "Flag for uploading File")
	Rootcmd.AddCommand(postCmd)
}

//  -H "content-Type":"json/content-Type","key2":"value2","key3":"value2", "key4":"value2"

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	Rootcmd.AddCommand(postCmd)
}

//  -H "content-Type":"json/content-Type","key2":"value2","key3":"value2", "key4":"value2"

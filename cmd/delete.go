package cmd

import (
	// "encoding/json"
	"fmt"
	"io"
	"pmc/internal/request"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "DELETE {url}",
	Short: "Send a HTTP DELETE request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body, _ := cmd.Flags().GetString("body")
		// header, _ := cmd.Flags().GetString("header")
		// var headerData map[string]string
		// json.Unmarshal([]byte(header), &headerData)
		fmt.Printf("Sending DELETE request to: %s\n", url)
		resp, err := request.RequestHandler("DELETE", url, strings.NewReader(body), nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error Reading Response ")
			return
		}

		fmt.Printf("Response: %s\n", string(data))
	},
}

func init() {
	deleteCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the delete request")
	deleteCmd.Flags().StringP("header", "H", "", "Provide custom headers in the format key:value,key:value")
	Rootcmd.AddCommand(deleteCmd)
}

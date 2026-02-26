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

var putCmd = &cobra.Command{
	Use:   "PUT {url}",
	Short: "Send a HTTP PUT request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		var headerData map[string]string
		err := json.Unmarshal([]byte(header), &headerData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sending PUT request to: %s\n", url)
		resp, err := request.RequestHandler("PUT", url, strings.NewReader(body), headerData)
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error Reading Response")
			return
		}
		fmt.Printf("Response: %s\n", string(data))
	},
}

func init() {
	putCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the put request")
	putCmd.Flags().StringP("header", "H", "", "Provide any header info")
	Rootcmd.AddCommand(putCmd)
}

package cmd

import (
	"fmt"
	"io"
	"pmc/internal/request"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "GET {url}",
	Short: "Send a HTTP GET request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending GET request to: %s\n", url)
		resp, err := request.RequestHandler("GET", url, nil, nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("Error Reading response")
			return
		}
		data, err := io.ReadAll(resp.Body)
		fmt.Printf("Response: %s\n", string(data))
	},
}

func init() {
	Rootcmd.AddCommand(getCmd)
}

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

var patchCmd = &cobra.Command{
	Use:   "PATCH [url]",
	Short: "Send a HTTP PATCH request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending PATCH request to: %s\n", url)
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		var headerData map[string]string
		err := json.Unmarshal([]byte(header), &headerData)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := request.RequestHandler("PATCH", url, strings.NewReader(body), headerData)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
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
	patchCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the Patch request")
	patchCmd.Flags().StringP("header", "H", "", "Provide Header info")
	Rootcmd.AddCommand(patchCmd)
}

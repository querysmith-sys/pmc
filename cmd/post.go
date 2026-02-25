package cmd

import (
	"fmt"
	"io"
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
		// fmt.Printf("Sending POST request to: %s\n", url)
		body, _ := cmd.Flags().GetString("body")
		resp, err := request.RequestHandler("POST", url, strings.NewReader(body))

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
	Rootcmd.AddCommand(postCmd)
}

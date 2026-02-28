package cmd

import (
	"fmt"
	"io"
	bodybuilder "pmc/internal/body-builder"
	"pmc/internal/request"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "PUT {url}",
	Short: "Send a HTTP PUT request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending PUT request to: %s\n", url)
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		filePath, _ := cmd.Flags().GetString("file")

		bodyReader, headerData := bodybuilder.Bodybuilder(filePath, header, body)
		resp, err := request.RequestHandler("PUT", url, bodyReader, headerData)
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
	putCmd.Flags().StringP("file", "F", "", "Flag for uploading File")
	Rootcmd.AddCommand(putCmd)
}

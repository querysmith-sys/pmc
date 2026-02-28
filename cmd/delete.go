package cmd

import (
	"fmt"
	"io"
	bodybuilder "pmc/internal/body-builder"
	"pmc/internal/request"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "DELETE {url}",
	Short: "Send a HTTP DELETE request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		filePath, _ := cmd.Flags().GetString("file")

		fmt.Printf("Sending DELETE request to: %s\n", url)

		bodyReader, headerData := bodybuilder.Bodybuilder(filePath, header, body)
		resp, err := request.RequestHandler("DELETE", url, bodyReader, headerData)
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
	deleteCmd.Flags().StringP("file", "F", "", "Flag for uploading File")
	Rootcmd.AddCommand(deleteCmd)
}

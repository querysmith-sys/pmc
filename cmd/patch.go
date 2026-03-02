package cmd

import (
	"fmt"
	bodybuilder "pmc/internal/body-builder"
	"pmc/internal/executor"

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

		filePath, _ := cmd.Flags().GetString("file")

		bodyReader, headerData := bodybuilder.Bodybuilder(filePath, header, body)

		executor.Executor("PATCH", url, bodyReader, headerData)
	},
}

func init() {
	patchCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the Patch request")
	patchCmd.Flags().StringP("file", "F", "", "Flag for uploading File")
	patchCmd.Flags().StringP("header", "H", "", "Provide Header info")
	Rootcmd.AddCommand(patchCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "PATCH {url}",
	Short: "Send a HTTP PATCH request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending PATCH request to: %s\n", url)
	},
}

func init() {
	patchCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the Patch request")
	Rootcmd.AddCommand(patchCmd)
}

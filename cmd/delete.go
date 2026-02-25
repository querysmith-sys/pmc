package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "DELETE {url}",
	Short: "Send a HTTP DELETE request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending DELETE request to: %s\n", url)
	},
}

func init() {
	deleteCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the delete request")
	Rootcmd.AddCommand(deleteCmd)
}

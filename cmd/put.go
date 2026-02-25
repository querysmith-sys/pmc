package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "PUT {url}",
	Short: "Send a HTTP PUT request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending PUT request to: %s\n", url)
	},
}

func init() {
	putCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the put request")
	Rootcmd.AddCommand(putCmd)
}

package cmd

import (
	"fmt"
	"pmc/internal/executor"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "GET {url}",
	Short: "Send a HTTP GET request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending GET request to: %s\n", url)

		executor.Executor("GET", url, nil, nil)
	},
}

func init() {
	Rootcmd.AddCommand(getCmd)
}

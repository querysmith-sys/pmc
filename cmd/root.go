package cmd

import (
	"github.com/spf13/cobra"
)

var Rootcmd = &cobra.Command{
	Use:   "httpreq",
	Short: "A simple and easy to understand http client cli. Alternate to postman in http cli version.",
}

func Execute() {
	Rootcmd.Execute()
}

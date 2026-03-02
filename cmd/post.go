package cmd

import (
	"fmt"
	"os"
	bodybuilder "pmc/internal/body-builder"
	"pmc/internal/executor"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "POST [url]",
	Short: "Send a HTTP POST request.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Sending POST request to: %s\n", url)
		body, _ := cmd.Flags().GetString("body")
		header, _ := cmd.Flags().GetString("header")
		filePath, _ := cmd.Flags().GetString("file")

		bodyReader, headerData := bodybuilder.Bodybuilder(filePath, header, body)

		body, status, headerDetails, responsemethod, contentLength, statusCode, timeTaken := executor.Executor("POST", url, bodyReader, headerData)
		//check for verbose to be true
		showExtraDetails, _ := cmd.Flags().GetBool("verbose")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if showExtraDetails == true {

			fmt.Fprintf(w, "Method:\t%s\n", responsemethod)
			fmt.Fprintf(w, "Content-Length:\t%d\n", contentLength)
			fmt.Fprintf(w, "Headers:\t%v\n", headerDetails)
		}

		fmt.Fprintln(w, "HTTP RESPONSE")
		fmt.Fprintf(w, "Status Code:\t%d (%s)\n", statusCode, status)
		fmt.Fprintf(w, "Time Taken:\t%s\n", timeTaken)
		fmt.Fprintf(w, "Body:\t%s\n", body)

		w.Flush()
	},
}

func init() {
	// create a flag for body
	postCmd.Flags().StringP("body", "b", "", "Provide Data to be send in the body of the post request")
	// defineing a  flag for custom header in the format of key and value and mulyiple header can be added by seperating them with comma and key and value should be seperated by colon
	postCmd.Flags().StringP("header", "H", "", "Provide custom headers in the format key:value,key:value")
	// defining a flag for  multipart/formData file upload
	postCmd.Flags().StringP("file", "F", "", "Flag for uploading File")
	// this flag is vervose to see extra info like the headers
	postCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	Rootcmd.AddCommand(postCmd)
}

//  -H "content-Type":"json/content-Type","key2":"value2","key3":"value2", "key4":"value2"

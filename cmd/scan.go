package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"engine/service"
)

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute the given inputs and generate a report",
	Long: `Generate a report with the given input file as describe by:

	{
		"id": "ABCD-1234",
		"content": [
			"QXJjA...."
		]
	}
`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		var rawInfo []byte
		var err error
		if input == "-" {
			rawInfo, err = io.ReadAll(os.Stdin)
		} else {
			rawInfo, err = os.ReadFile(input)
		}

		if err != nil {
			errorAndClose(err, output)
		}

		var request service.Request
		err = json.Unmarshal(rawInfo, &request)
		if err != nil {
			errorAndClose(err, output)
		}

		compute := service.NewCompute()
		report, err := compute.Process(request)
		if err != nil {
			errorAndClose(err, output)
		}

		payload, err := json.MarshalIndent(report, "", "    ")
		if err != nil {
			errorAndClose(err, output)
		}

		// Output to STDOUT
		if output == "-" {
			fmt.Println(string(payload))
			return
		}

		f, err := os.Create(fmt.Sprintf("%s.json", output))
		if err != nil {
			errorAndClose(err, output)
		}
		defer f.Close()

		_, err = f.Write(payload)
		if err != nil {
			errorAndClose(err, output)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(computeCmd)
	computeCmd.Flags().StringP("input", "i", "-", "Path of the input file, if value is '-' then read from STDIN")
	computeCmd.Flags().StringP("output", "o", "-", "Path of the output file without extension, if value is '-' then write to STDOUT")
}

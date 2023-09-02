package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var licencecmd = &cobra.Command{
	Use:   "licence",
	Short: "Read licence",
	Long:  `Read licence`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("This is a cli tool that is used to trace peoples ip infon from the internate which is avelabel for public to see pleas use it In a wise manner")
		}
	},
}

func init() {
	rootCmd.AddCommand(licencecmd)

}

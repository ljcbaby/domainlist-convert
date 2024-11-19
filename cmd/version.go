package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "-dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of p2d",
	Long:  `All software has versions. This is p2d's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("p2d v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"os"

	"github.com/dn-11/provider2domainset/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"
)

var rootCmd = &cobra.Command{
	Use:   "provider2domainset",
	Short: "p2d, convert domain list form clash provider to mosdns domainset.",
	Long:  `p2d, convert domain list form clash provider to mosdns domainset.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			log.SetLogLevel(zapcore.DebugLevel)
		}
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
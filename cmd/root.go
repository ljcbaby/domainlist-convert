package cmd

import (
	"os"

	"github.com/ljcbaby/domainlist-convert/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"
)

var (
	config string
)

var rootCmd = &cobra.Command{
	Use:   "provider2domainset",
	Short: "p2d, convert domain list form clash provider to mosdns domainset.",
	Long:  `p2d, convert domain list form clash provider to mosdns domainset.`,
}

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
			log.L().SetLogLevel(zapcore.DebugLevel)
		}
	}
}

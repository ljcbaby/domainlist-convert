package cmd

import (
	"github.com/ljcbaby/domainlist-convert/conf"
	"github.com/ljcbaby/domainlist-convert/convert"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run convert once",
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init(config)
		convert.RunOnce()
	},
}

func init() {
	runCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "config file path")
	rootCmd.AddCommand(runCmd)
}

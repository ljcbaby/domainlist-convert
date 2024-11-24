package cmd

import (
	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/convert"
	"github.com/dn-11/provider2domainset/daemon"
	"github.com/dn-11/provider2domainset/log"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "run as service",
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init(config)
		if conf.Service.Enable {
			daemon.Serve()
		} else {
			log.L().Warn("Service not enabled, run once.")
			convert.RunOnce()
		}
	},
}

func init() {
	serviceCmd.PersistentFlags().StringVarP(&config, "config", "c", "/etc/provider2domainset.yaml", "config file path")
	rootCmd.AddCommand(serviceCmd)
}

package cmd

import (
	"github.com/dn-11/provider2domainset/daemon"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install provider2domainset to /usr/sbin/provider2domainset",
	Run: func(cmd *cobra.Command, args []string) {
		daemon.Install()
		daemon.AddService()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

package cmd

import (
	"github.com/ljcbaby/domainlist-convert/daemon"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstallation command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall provider2domainset from /usr/sbin/provider2domainset",
	Run: func(cmd *cobra.Command, args []string) {
		daemon.RmService()
		daemon.Uninstall()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

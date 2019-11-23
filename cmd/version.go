package cmd

import (
	"github.com/spf13/cobra"
)

// Go Releaser Variables
var builtBy, commit, date, version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output version info",
	Long:  `Output version info for this tool as well as build date and commit hash.`,
	Run: func(cmd *cobra.Command, args []string) {
		status("Version	", version)
		status("Commit	", commit)
		status("Date	", date)
		status("BuiltBy	", builtBy)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

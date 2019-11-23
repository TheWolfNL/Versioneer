package cmd

import (
	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Release a Minor version",
	Long:  `Create a tag for a Minor version higher than latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		var releaseType releaseType = Minor
		release(releaseType)
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)
}

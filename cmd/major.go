package cmd

import (
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Release a Major version",
	Long:  `Create a tag for a Major version higher than latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		var releaseType releaseType = Major
		release(releaseType)
	},
}

func init() {
	rootCmd.AddCommand(majorCmd)
}

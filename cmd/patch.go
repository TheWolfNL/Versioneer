package cmd

import (
	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Release a Patch version",
	Long:  `Create a tag for a Patch version higher than latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		var releaseType releaseType = Patch
		release(releaseType)
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var applicationName = "Versioneer"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   strings.ToLower(applicationName),
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
  examples and usage of using your application. For example:
  
  Cobra is a CLI library for Go that empowers applications.
  This application is a tool to generate the needed files
  to quickly create a Cobra application.`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func release(releaseType releaseType) {
	latest := getLatestVersion()
	status("Release Type", yellow(releaseType.String()))
	status("Previous Version", yellow(latest))
	newVersion := incrementVersion(releaseType)
	status("Version", yellow(newVersion))
	//TODO: add confirm
	/*
	  if is_CI || confirm "Do you want to release '$(Blue $(tag))$(Yellow \'?) [Y/n]"; then
	*/
	output(fmt.Sprintf("Releasing version '%s'", yellow(newVersion)))
	// createAndPushTag(newVersion)
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.PersistentFlags().StringVar(&accessToken, "access-token", accessToken,
		"A personal GitHub access token for API authentication. "+
			"This option overrides the GITHUB_ACCESS_TOKEN environment variable.")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", timeout,
		"Timeout to set for GitHub API calls.")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", verbose,
		"Output entire license JSON response.")
}

var rootCmd = &cobra.Command{
	Use:   "ghlicense",
	Short: "A lookup tool for GitHub project licenses.",
	Long: `A lookup tool for GitHub project licenses.

Logging is controlled via these environment variables:
- LOG_LEVEL: debug, info, warn, error, fatal
- LOG_COLOR: true, false
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

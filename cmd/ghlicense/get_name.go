package main

import (
	"context"
	"fmt"

	"github.com/petergardfjall/ghlicense/pkg/gh"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getNameCmd)
}

var getNameCmd = &cobra.Command{
	Use:   "name <repo-url>",
	Short: "Get the GitHub-reported license of a repo.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repoURL := args[0]

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		client := gh.NewClient(ctx, mustGetAccessToken())
		lic, err := client.GetLicense(ctx, repoURL)
		if err != nil {
			return errors.Wrap(err, "GetLicense call failed")
		}

		if verbose {
			log.Debug().Msg(prettyJSON(lic))
		}

		fmt.Println(*lic.License.SPDXID)

		return nil
	},
}

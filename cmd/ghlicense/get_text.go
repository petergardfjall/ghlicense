package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/petergardfjall/ghlicense/pkg/gh"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getTextCmd)
}

var getTextCmd = &cobra.Command{
	Use:   "text <repo-url>",
	Short: "Get license text for the GitHub-reported license of a repo.",
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

		decoded, err := fromBase64(*lic.Content)
		if err != nil {
			return errors.Wrap(err, "base64-decoding failed")
		}

		fmt.Println(decoded)

		return nil
	},
}

func fromBase64(encoded string) (string, error) {
	src := []byte(encoded)
	dest := make([]byte, len(src))
	if _, err := base64.StdEncoding.Decode(dest, src); err != nil {
		return "", err
	}
	return string(dest), nil
}
